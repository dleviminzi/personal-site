package pages

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	site "github.com/dleviminzi/site"
	_ "github.com/lib/pq"
)

// ProjectList contains array of Project
type ProjectList struct {
	l         *log.Logger
	db        *sql.DB
	Projects  []site.Project
	PageTitle string
}

// HandleProjectsList constructs a new list of
func NewProjectsList(l *log.Logger, db *sql.DB) *ProjectList {
	return &ProjectList{l, db, []site.Project{}, "Projects"}
}

func (pl *ProjectList) dbFetch() error {
	var fetchedProjects []site.Project

	// query the database
	rows, err := pl.db.Query(site.ProjectListQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// iterate through result rows and return array
	for rows.Next() {
		var project site.Project

		if err = rows.Scan(&project.Title, &project.Description, &project.GithubLink,
			&project.Status, &project.StartDate, &project.EndDate); err != nil {
			return err
		}

		fetchedProjects = append(fetchedProjects, project)
	}

	if err = rows.Err(); err != nil {
		pl.Projects = fetchedProjects
		return err
	}
	pl.Projects = fetchedProjects
	return nil
}

func (pl *ProjectList) Serve(w http.ResponseWriter, r *http.Request) {
	err := pl.dbFetch()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not locate that note"))
		return
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// set up template and write it
	projectTemp := template.Must(template.New("projects").ParseFiles(htmlTemplates("projects")...))
	err = projectTemp.Execute(w, pl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to load template for projects"))
		return
	}
}

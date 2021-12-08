package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dleviminzi/personal-site/data"
	_ "github.com/lib/pq"
)

// ProjectList contains array of Project
type ProjectList struct {
	l         *log.Logger
	db        *sql.DB
	Projects  []data.Project
	PageTitle string
}

// HandleProjectsList constructs a new list of
func HandleProjectsList(l *log.Logger, db *sql.DB) *ProjectList {
	return &ProjectList{l, db, []data.Project{}, "Projects"}
}

func (pl *ProjectList) dbFetch() error {
	var fetchedProjects []data.Project

	// query the database
	rows, err := pl.db.Query(data.ProjectListQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// iterate through result rows and return array
	for rows.Next() {
		var project data.Project

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

func (pl *ProjectList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := pl.dbFetch()
	if err != nil {
		pl.l.Fatal(err)
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// set up template and write it
	projectTemp := template.Must(template.New("projects").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/projects.html", "./templates/footer.html"))
	err = projectTemp.Execute(w, pl)
	if err != nil {
		pl.l.Fatal(err)
	}
}

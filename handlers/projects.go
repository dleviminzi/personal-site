package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dleviminzi/personal-site/data"
	_ "github.com/lib/pq"
)

/*

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

*/

// ProjectList contains array of Project
type ProjectList struct {
	logger    *log.Logger
	db        *sql.DB
	Projects  []data.Project
	PageTitle string
}

// NewProjectsList constructs a new list of
func NewProjectsList(l *log.Logger, db *sql.DB) *ProjectList {
	return &ProjectList{l, db, []data.Project{}, "Projects"}
}

func (projectList *ProjectList) dbFetch() error {
	var fetchedProjects []data.Project

	// query the database
	rows, err := projectList.db.Query(data.ProjectListQuery)
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
		projectList.Projects = fetchedProjects
		return err
	}
	projectList.Projects = fetchedProjects
	return nil
}

func (pl *ProjectList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := pl.dbFetch()
	if err != nil {
		pl.logger.Fatal(err)
	}

	// set up template and write it
	projectTemplate := template.Must(template.New("projects").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/projects.html", "./templates/footer.html"))
	err = projectTemplate.Execute(w, pl)
	if err != nil {
		pl.logger.Fatal(err)
	}
}

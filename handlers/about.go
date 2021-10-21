package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dleviminzi/personal-site/data"
)

// About simply contains a logger to send welcome message through
type About struct {
	logger      *log.Logger
	db          *sql.DB
	Experiences []data.ExperienceItem
	PageTitle   string
}

// NewAbout constructs a Welcome
func NewAbout(logger *log.Logger, db *sql.DB) *About {
	return &About{logger, db, []data.ExperienceItem{}, "About"}
}

func (about *About) dbFetch() error {
	var fetchedExperiences []data.ExperienceItem

	// query the database
	rows, err := about.db.Query(data.ExperienceItemsQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// iterate through result rows and return array
	for rows.Next() {
		var exp data.ExperienceItem

		if err = rows.Scan(&exp.ItemType, &exp.Title, &exp.Description, &exp.StartDate, &exp.EndDate); err != nil {
			return err
		}

		fetchedExperiences = append(fetchedExperiences, exp)
	}

	if err = rows.Err(); err != nil {
		about.Experiences = fetchedExperiences
		return err
	}
	about.Experiences = fetchedExperiences
	return nil
}

func (about *About) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := about.dbFetch()
	if err != nil {
		about.logger.Fatal(err)
	}

	// set up template and write it
	aboutTemplate := template.Must(template.New("about").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/about.html", "./templates/footer.html"))
	err = aboutTemplate.Execute(responseWriter, about)
	if err != nil {
		about.logger.Fatal(err)
	}
}

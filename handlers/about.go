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
	l           *log.Logger
	db          *sql.DB
	Experiences []data.ExperienceItem
	PageTitle   string
}

// HandleAbout constructs a Welcome
func HandleAbout(l *log.Logger, db *sql.DB) *About {
	return &About{l, db, []data.ExperienceItem{}, "About"}
}

func (a *About) dbFetch() error {
	var fetchedExperiences []data.ExperienceItem

	// query the database
	rows, err := a.db.Query(data.ExperienceItemsQuery)
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
		a.Experiences = fetchedExperiences
		return err
	}
	a.Experiences = fetchedExperiences
	return nil
}

func (a *About) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := a.dbFetch()
	if err != nil {
		a.l.Fatal(err)
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// set up template and write it
	aboutTemp := template.Must(template.New("about").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/about.html", "./templates/footer.html"))
	err = aboutTemp.Execute(w, a)
	if err != nil {
		a.l.Fatal(err)
	}
}

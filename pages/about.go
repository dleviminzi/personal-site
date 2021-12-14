package pages

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	site "github.com/dleviminzi/site"
)

// About is just the index page of the site with my bio
type About struct {
	l           *log.Logger
	db          *sql.DB
	Experiences []site.ExperienceItem
	PageTitle   string
}

// HandleAbout constructs a Welcome
func NewAbout(l *log.Logger, db *sql.DB) *About {
	return &About{l, db, []site.ExperienceItem{}, "about"}
}

func (a *About) dbFetch() error {
	var fetchedExperiences []site.ExperienceItem

	// query the database
	rows, err := a.db.Query(site.ExperienceItemsQuery)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Failed to close rows... Probably not a good sign")
		}
	}(rows)

	// iterate through result rows and return array
	for rows.Next() {
		var exp site.ExperienceItem

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

func (a *About) Serve(w http.ResponseWriter, r *http.Request) {
	err := a.dbFetch()
	if err != nil {
		a.l.Fatal(err)
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// set up template and write it
	aboutTemp := template.Must(template.New("about").ParseFiles(htmlTemplates("about")...))
	err = aboutTemp.Execute(w, a)
	if err != nil {
		a.l.Fatal(err)
	}
}

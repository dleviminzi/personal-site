package pages

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/dleviminzi/site"
)

// NoteList contains array of Project
type NoteList struct {
	l         *log.Logger
	db        *sql.DB
	Notes     []site.Note
	PageTitle string
}

// HandleNotesList constructs a new list of
func NewNotesList(l *log.Logger, db *sql.DB) *NoteList {
	return &NoteList{l, db, []site.Note{}, "Notes"}
}

func (nl *NoteList) dbFetch() error {
	var fetchedNotes []site.Note

	// query the database
	rows, err := nl.db.Query(site.NoteListQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// TODO: perhaps add FromDB method to struct and FromJSON method as well
	// would make things cleaner
	for rows.Next() {
		var note site.Note

		if err = rows.Scan(&note.Title, &note.Topic, &note.Content); err != nil {
			return err
		}

		fetchedNotes = append(fetchedNotes, note)
	}

	if err = rows.Err(); err != nil {
		nl.Notes = fetchedNotes
		return err
	}
	nl.Notes = fetchedNotes
	return nil
}

func (nl *NoteList) Serve(w http.ResponseWriter, r *http.Request) {
	err := nl.dbFetch()
	if err != nil {
		nl.l.Fatal(err)
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// set up template and write it
	// TODO: update templates directory to html directory and move that directory into this folder!
	nlTemp := template.Must(template.New("notes").ParseFiles(htmlTemplates("notes")...))
	err = nlTemp.Execute(w, nl)
	if err != nil {
		nl.l.Fatal(err)
	}
}

package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dleviminzi/personal-site/data"
	_ "github.com/lib/pq"
)

// NoteList contains array of Project
type NoteList struct {
	l         *log.Logger
	db        *sql.DB
	Notes     []data.Note
	PageTitle string
}

// HandleNotesList constructs a new list of
func HandleNotesList(l *log.Logger, db *sql.DB) *NoteList {
	return &NoteList{l, db, []data.Note{}, "Notes"}
}

func (nl *NoteList) dbFetch() error {
	var fetchedNotes []data.Note

	// query the database
	rows, err := nl.db.Query(data.NoteListQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// TODO: perhaps add FromDB method to struct and FromJSON method as well
	// would make things cleaner
	for rows.Next() {
		var note data.Note

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

func (nl *NoteList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := nl.dbFetch()
	if err != nil {
		nl.l.Fatal(err)
	}
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// set up template and write it
	nlTemp := template.Must(template.New("notes").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/notes.html", "./templates/footer.html"))
	err = nlTemp.Execute(w, nl)
	if err != nil {
		nl.l.Fatal(err)
	}
}

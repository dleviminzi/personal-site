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
	logger    *log.Logger
	db        *sql.DB
	Notes     []data.Note
	PageTitle string
}

// NewNotesList constructs a new list of
func NewNotesList(logger *log.Logger, db *sql.DB) *NoteList {
	return &NoteList{logger, db, []data.Note{}, "Notes"}
}

func (noteList *NoteList) dbFetch() error {
	var fetchedNotes []data.Note

	// query the database
	rows, err := noteList.db.Query(data.NoteListQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	// iterate through result rows and return array
	for rows.Next() {
		var note data.Note

		if err = rows.Scan(&note.Title, &note.Topic, &note.Content); err != nil {
			return err
		}

		fetchedNotes = append(fetchedNotes, note)
	}

	if err = rows.Err(); err != nil {
		noteList.Notes = fetchedNotes
		return err
	}
	noteList.Notes = fetchedNotes
	return nil
}

func (noteList *NoteList) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := noteList.dbFetch()
	if err != nil {
		noteList.logger.Fatal(err)
	}

	// set up template and write it
	noteListTemplate := template.Must(template.New("notes").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/notes.html", "./templates/footer.html"))
	err = noteListTemplate.Execute(responseWriter, noteList)
	if err != nil {
		noteList.logger.Fatal(err)
	}
}

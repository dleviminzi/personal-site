package pages

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dleviminzi/site"
	"github.com/gorilla/mux"
)

type Note struct {
	l         *log.Logger
	db        *sql.DB
	PageTitle string
	NoteTitle string
	Content   site.Note
}

func NewNote(l *log.Logger, db *sql.DB) *Note {
	return &Note{l, db, "Notes", "", site.Note{}}
}

func (n *Note) dbFetch(title string) error {
	err := n.db.QueryRow(site.NoteQuery, title).Scan(&n.Content.Title, &n.Content.Topic, &n.Content.Content)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Serve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n.NoteTitle = vars["noteName"]

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := n.dbFetch(n.NoteTitle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not locate that note"))
		return
	}

	// set up template and write it
	noteTemp := template.Must(template.New("note").ParseFiles(htmlTemplates("note")...))
	err = noteTemp.Execute(w, n) // we pass the writer as output and the note for content
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to load template for the note"))
		return
	}
}

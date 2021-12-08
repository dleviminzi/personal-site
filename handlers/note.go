package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dleviminzi/personal-site/data"
	"github.com/gorilla/mux"
)

type ReqNote struct {
	l         *log.Logger
	db        *sql.DB
	PageTitle string
	NoteTitle string
	ReqNote   data.Note
}

func HandleNote(l *log.Logger, db *sql.DB) *ReqNote {
	return &ReqNote{l, db, "Notes", "", data.Note{}}
}

func (n *ReqNote) dbFetch(title string) error {
	err := n.db.QueryRow(data.NoteQuery, title).Scan(&n.ReqNote.Title, &n.ReqNote.Topic, &n.ReqNote.Content)
	if err != nil {
		return err
	}
	return nil
}

func (n *ReqNote) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n.NoteTitle = vars["noteName"]

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := n.dbFetch(n.NoteTitle)
	if err != nil {
		n.l.Fatal(err)
	}

	// set up template and write it
	noteTemp := template.Must(template.New("note").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/note.html", "./templates/footer.html"))
	err = noteTemp.Execute(w, n)
	if err != nil {
		n.l.Fatal(err)
	}
}

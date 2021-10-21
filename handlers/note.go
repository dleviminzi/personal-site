package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/dleviminzi/personal-site/data"
)

type ReqNote struct {
	logger    *log.Logger
	db        *sql.DB
	PageTitle string
	NoteTitle string
	ReqNote   data.Note /* this name is bad */
}

func NewNote(logger *log.Logger, db *sql.DB, path string) *ReqNote {
	// do some stuff to get note id from path
	pathSegs := strings.Split(path, "/")
	title := pathSegs[len(pathSegs)-1]

	return &ReqNote{logger, db, "Notes", title, data.Note{}}
}

func (note *ReqNote) dbFetch(title string) error {
	err := note.db.QueryRow(data.NoteQuery, title).Scan(&note.ReqNote.Title, &note.ReqNote.Topic, &note.ReqNote.Content)
	if err != nil {
		return err
	}
	return nil
}

func (note *ReqNote) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	err := note.dbFetch(note.NoteTitle)
	if err != nil {
		note.logger.Fatal(err)
	}

	// set up template and write it
	noteTemplate := template.Must(template.New("note").ParseFiles("./templates/header.html", "./templates/nav.html", "./templates/note.html", "./templates/footer.html"))
	err = noteTemplate.Execute(responseWriter, note)
	if err != nil {
		note.logger.Fatal(err)
	}
}

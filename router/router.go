package router

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"

	"github.com/dleviminzi/site/pages"
)

//
// Router for my sites pages
//
func New(l *log.Logger, db *sql.DB) *mux.Router {
	noteList := pages.NewNotesList(l, db)
	note := pages.NewNote(l, db)
	about := pages.NewAbout(l, db)
	projectList := pages.NewProjectsList(l, db)

	r := mux.NewRouter()
	notesSR := r.PathPrefix("/notes").Subrouter() /* we have note list and specific notes */
	notesSR.HandleFunc("/", noteList.Serve).Methods("GET")
	notesSR.HandleFunc("/{noteName}", note.Serve).Methods("GET")
	r.HandleFunc("/", about.Serve).Methods("GET")
	r.HandleFunc("/about", about.Serve).Methods("GET")
	r.HandleFunc("/projects/", projectList.Serve).Methods("GET")
	return r
}

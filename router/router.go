package router

import (
	"database/sql"
	"log"

	"github.com/dleviminzi/personal-site/handlers"
	"github.com/gorilla/mux"
)

//
// Router for my sites pages
//
func New(l *log.Logger, db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	notesSR := r.PathPrefix("/notes").Subrouter() /* we have note list and specific notes */
	notesSR.HandleFunc("/", handlers.HandleNotesList(l, db).ServeHTTP).Methods("GET")
	notesSR.HandleFunc("/{noteName}", handlers.HandleNote(l, db).ServeHTTP).Methods("GET")
	r.HandleFunc("/", handlers.HandleAbout(l, db).ServeHTTP).Methods("GET")
	r.HandleFunc("/about", handlers.HandleAbout(l, db).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects/", handlers.HandleProjectsList(l, db).ServeHTTP).Methods("GET")
	return r
}

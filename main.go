package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/dleviminzi/personal-site/handlers"
)

func main() {
	logger := log.New(os.Stdout, "api", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		logger.Fatal(err)
	}

	// Route handling logic
	r := mux.NewRouter()
	notesSR := r.PathPrefix("/notes").Subrouter() /* we have note list and specific notes */
	notesSR.HandleFunc("/", handlers.HandleNotesList(logger, db).ServeHTTP).Methods("GET")
	notesSR.HandleFunc("/{noteName}", handlers.HandleNote(logger, db).ServeHTTP).Methods("GET")
	r.HandleFunc("/", handlers.HandleAbout(logger, db).ServeHTTP).Methods("GET")
	r.HandleFunc("/about", handlers.HandleAbout(logger, db).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects/", handlers.HandleProjectsList(logger, db).ServeHTTP).Methods("GET")

	server := NewServer(r, ":9090")

	go func() {
		err := server.ListenAndServeTLS("cert.pem", "key.pem")
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Set up a channel for interupt/terminate signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)

	// Block waiting for interupt/terminate input
	sig := <-sigChan
	logger.Println("Received Terminate, Shutting Down.", sig)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutContext); err != nil {
		logger.Fatal(err)
	}
}

func NewServer(router http.Handler, address string) *http.Server {
	s := &http.Server{
		Addr:         address,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	return s
}

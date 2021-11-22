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
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.NewAbout(logger, db).ServeHTTP)
	router.HandleFunc("/about", handlers.NewAbout(logger, db).ServeHTTP)
	router.HandleFunc("/projects/", handlers.NewProjectsList(logger, db).ServeHTTP)
	noteSubrouter := router.PathPrefix("/notes").Subrouter() /* we have note list and specific notes */
	noteSubrouter.HandleFunc("/", handlers.NewNotesList(logger, db).ServeHTTP)
	noteSubrouter.HandleFunc("/{noteName:[a-z]+}/", handlers.NewNote(logger, db).ServeHTTP)

	server := &http.Server{
		Addr:         ":9990",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
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

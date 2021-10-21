package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dleviminzi/personal-site/handlers"
)

func main() {
	logger := log.New(os.Stdout, "api", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		logger.Fatal(err)
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", Serve(logger, db))

	server := &http.Server{
		Addr:         ":9990",
		Handler:      serveMux,
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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received Terminate, Shutting Down.", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

func match(givenPath string, existingPath string) bool {
	// TODO: figure out how to handle error here
	existingPathRegExp := regexp.MustCompile(existingPath)
	result := existingPathRegExp.MatchString(givenPath)

	return result
}

// Serve takes the path and determines the proper handler
// this is largely based on https://benhoyt.com/writings/go-routing/
func Serve(logger *log.Logger, db *sql.DB) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var handler http.Handler
		path := request.URL.Path

		switch {
		case match(path, `^/$`): /* return react app */
			handler = handlers.NewAbout(logger, db)
		case match(path, `^/about$`):
			handler = handlers.NewAbout(logger, db)
		case match(path, `^/projects[/]?$`):
			handler = handlers.NewProjectsList(logger, db)
		case match(path, `^/notes[/]?$`):
			handler = handlers.NewNotesList(logger, db)
		case match(path, `^/notes/([a-z]+)$`): /* return notes content for title */
			handler = handlers.NewNote(logger, db, path)
		default:
			http.NotFound(responseWriter, request)
			return
		}

		handler.ServeHTTP(responseWriter, request)
	}
}

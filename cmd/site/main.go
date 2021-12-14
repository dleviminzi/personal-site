package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dleviminzi/site/router"
	"github.com/dleviminzi/site/server"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: Errors should be handled less dramatically in all the packages
// TODO: Write tests

func main() {
	l := log.New(os.Stdout, "site: ", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		l.Fatal(err)
	}

	// Initiate router and server
	r := router.New(l, db)
	s := server.New(r, ":9090")

	go func() {
		err := s.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem")
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Set up a channel for interrupt/terminate signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)

	// Block waiting for interrupt/terminate input
	sig := <-sigChan
	l.Println("Received Terminate, Shutting Down.", sig)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(timeoutContext); err != nil {
		l.Fatal(err)
	}
}

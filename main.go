package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dleviminzi/personal-site/router"
	"github.com/dleviminzi/personal-site/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := log.New(os.Stdout, "site: ", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		logger.Fatal(err)
	}

	// Initiate router and server
	r := router.New(logger, db)
	server := server.New(r, ":9090")

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

package server

import (
	"net/http"
	"time"
)

//
// Server for the site
//
func New(router http.Handler, address string) *http.Server {
	s := &http.Server{
		Addr:         address,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	return s
}

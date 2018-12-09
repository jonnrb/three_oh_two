package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"go.jonnrb.io/three_oh_two/redir"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must provide redirect URL root")
	}
	h, err := redir.Handler(os.Args[1])
	if err != nil {
		log.Fatalf("Could not create redirect handler: %v", err)
	}

	log.Print((&http.Server{
		// This should be sufficient.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,

		Addr:    ":8080",
		Handler: h,
	}).ListenAndServe())
}

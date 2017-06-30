package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	goji "goji.io"
	"goji.io/pat"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func Pulse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT environment variable must be defined")
	}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), Hello)
	mux.HandleFunc(pat.Get("/pulse"), Pulse)

	log.Printf("Starting server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

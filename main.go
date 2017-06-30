package main

import (
	"fmt"
	"log"
	"net/http"

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
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), Hello)
	mux.HandleFunc(pat.Get("/pulse"), Pulse)

	log.Printf("Starting growser listening on :8000")
	http.ListenAndServe("localhost:8000", mux)
}

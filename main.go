/*
 * Copyright 2017 Thingful Ltd.

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	goji "goji.io"
	"goji.io/pat"
)

const (
	// DefaultPort is the default port to listen on should it not be supplied by
	// via the environment
	DefaultPort = "8080"
)

// Hello is a simple handler that extracts a name from the request url, and
// returns a string greeting to that name.
func Hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s", name)
}

// Pulse is a simple handler that returns ok when called. Used as a health
// check endpoint for our ELB.
func Pulse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), Hello)
	mux.HandleFunc(pat.Get("/pulse"), Pulse)

	log.Printf("Starting server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

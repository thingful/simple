/**
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

	// DefaultGreeting is used when generating our responses in the Hello handler
	DefaultGreeting = "Hello"
)

// Error is an interface we return from our custom handler types should any error occur
type Error interface {
	error
	Status() int
}

// HTTPError is an implementation of the Error interface above.
type HTTPError struct {
	Code int
	Err  error
}

// Error is an implementation of the Error interface for our custom type.
func (he HTTPError) Error() string {
	return he.Err.Error()
}

// Status returns the http status code of the current error.
func (he HTTPError) Status() int {
	return he.Code
}

// Env is a simple type for holding application specific context information.
// Here we just hold the text greeting string returned by Hello to demonstrate
// how this could be used.
type Env struct {
	Greeting string
}

// Handler is our custom HTTP handler type that implements the ServeHTTP method
// but is aware of the runtime Env context
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) Error
}

// ServeHTTP is our implementation of the http.Handler ServeHTTP function which
// makes our custom type into an http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// Hello is a simple handler that extracts a name from the request URL, and
// returns a string greeting to that name.
func Hello(env *Env, w http.ResponseWriter, r *http.Request) Error {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "%s, %s!", env.Greeting, name)
	return nil
}

// Pulse is a simple handler that returns ok when called. Used as a health
// check endpoint for our ELB.
func Pulse(env *Env, w http.ResponseWriter, r *http.Request) Error {
	fmt.Fprintf(w, "ok")
	return nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = DefaultGreeting
	}

	env := &Env{
		Greeting: greeting,
	}

	mux := goji.NewMux()
	mux.Handle(pat.Get("/hello/:name"), Handler{Env: env, H: Hello})
	mux.Handle(pat.Get("/pulse"), Handler{Env: env, H: Pulse})

	log.Printf("Starting server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

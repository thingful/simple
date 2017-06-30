package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	goji "goji.io"
	"goji.io/pat"
)

func TestHandlers(t *testing.T) {
	testcases := []struct {
		path     string
		status   int
		expected string
	}{
		{
			"/hello/world",
			http.StatusOK,
			"Hello, world!",
		},
		{
			"/pulse",
			http.StatusOK,
			"ok",
		},
	}

	for _, testcase := range testcases {
		req, err := http.NewRequest("GET", testcase.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := goji.NewMux()
		handler.HandleFunc(pat.Get("/hello/:name"), Hello)
		handler.HandleFunc(pat.Get("/pulse"), Pulse)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != testcase.status {
			t.Errorf("handler returned unexpected status code: got %v, wanted %v", status, testcase.status)
		}

		if rr.Body.String() != testcase.expected {
			t.Errorf("handler returned unexpected body: got %v, wanted %v", rr.Body.String(), testcase.expected)
		}
	}
}

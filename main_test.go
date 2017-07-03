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

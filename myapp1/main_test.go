package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// Set env variable
	os.Setenv("MESSAGE", "uwu from test")

	// Create a request to pass to the handler
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Use the handler from main.go
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("MESSAGE")))
	}).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if !strings.Contains(string(body), "uwu from test") {
		t.Errorf("expected MESSAGE to be in response, got %s", string(body))
	}
}

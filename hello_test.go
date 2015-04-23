package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDd struct{}

func TestHome(t *testing.T) {
	main := mainHandler()
	req, _ := http.NewRequest("GET", "", nil)

	w := httptest.NewRecorder()
	main.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page not avaliable, status: %v", w.Code)
	}
}

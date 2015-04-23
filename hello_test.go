package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type HandleTester func(method string, params url.Values) *httptest.ResponseRecorder

func GenerateHandleTester(t *testing.T, handleFunc http.Handler) HandleTester {
	return func(method string, params url.Values) *httptest.ResponseRecorder {
		req, err := http.NewRequest(method, "", strings.NewReader(params.Encode()))
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Set(
			"Content-Type",
			"application/x-www-form-urlencoded; param=value",
		)
		w := httptest.NewRecorder()
		handleFunc.ServeHTTP(w, req)
		return w
	}
}

func TestHome(t *testing.T) {
	homeHandle := mainHandler()
	test := GenerateHandleTester(t, homeHandle)
	w := test("GET", url.Values{})
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

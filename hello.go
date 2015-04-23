package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func mainHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Ansible and Travis! Route: %q", html.EscapeString(r.URL.Path))
	})
}

func main() {
	http.HandleFunc("/", mainHandler())
	log.Fatal(http.ListenAndServe(":80", nil))
}

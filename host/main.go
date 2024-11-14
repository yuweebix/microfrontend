package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/yuweebix/microfrontend/frontend/host/templates"
)

func main() {
	http.Handle("/", templ.Handler(templates.Index(), templ.WithContentType("text/html")))
	log.Println("Host server running on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

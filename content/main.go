package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/yuweebix/microfrontend/frontend/content/templates"
)

func main() {
	http.Handle("/", templ.Handler(templates.Table(), templ.WithContentType("text/html")))
	log.Println("Content server running on http://0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

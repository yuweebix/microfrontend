package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/yuweebix/microfrontend/frontend/host/templates"
)

func main() {
	r := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsHandler.Handler)
	r.Use(logClientIP)

	r.Handle("GET /", templ.Handler(templates.Index(), templ.WithContentType("text/html")))
	r.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		fmt.Println(string(data))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func logClientIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		fmt.Println("Client IP:", clientIP)
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/yuweebix/microfrontend/frontend/content/templates"
)

func main() {
	r := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsHandler.Handler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "close")

		f, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "data: ")
		templates.Content().Render(r.Context(), w)
		fmt.Fprintf(w, "\n\n")
		f.Flush()
		<-r.Context().Done()
	})

	log.Fatal(http.ListenAndServe(":8081", r))
}

package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// Enable Cors for the Frontend
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", " DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Post("/image", HandleIncomingImage)

	r.Post("/mattermost_callback", HandleCallbackFromMattermost)

	log.Println("Mattermost-Connector started")
	http.ListenAndServe(":80", r)
}

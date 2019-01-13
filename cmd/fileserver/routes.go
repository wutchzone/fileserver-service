package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func initRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(handleLog)
		r.Get("/", handleHello)
		r.Route("/api", func(r chi.Router) {
			r.Get("/", HandleGetAll)
			r.Get("/{name}", HandleGetOne)
		})
	})

	return r
}

func handleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Server is listening on port: %d", PORT)))
}

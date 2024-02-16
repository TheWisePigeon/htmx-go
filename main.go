package main

import (
	"htmx-go/templates"
	"log"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Dashboard.ExecuteWriter(pongo2.Context{}, w)
		return
	})
  log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bradj/iridium/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const configLocation string = "config.toml"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprintln(w, "home")
		return
	}

	tmpl.Execute(w, nil)
}

func main() {
	//c, err := NewConfig(configLocation)
	port := 3000

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)

	r.Group(routes.Protected)
	r.Group(routes.Public)

	log.Printf("Listening on %d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

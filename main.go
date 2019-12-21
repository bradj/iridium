package main

import (
	"fmt"
	"html/template"
	"net/http"

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
	// c, err := NewConfig(configLocation)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)
	r.Get("/login", loginHandler)
	r.Get("/logout", logoutHandler)
	r.Post("/upload", uploadHandler)

	http.ListenAndServe(":3000", r)
}

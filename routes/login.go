package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/go-chi/chi"
)

func (h HTTP) loginMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.loginGet)
	r.Post("/", h.loginPost)

	return r
}

func (h HTTP) loginGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")

	if err != nil {
		fmt.Fprintln(w, "login")
		return
	}

	tmpl.Execute(w, nil)
}

func (h HTTP) loginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	h.App.Logger.Printf("login request for user '%v'", username)

	if username != "brad" || password != "mypass" {
		fmt.Fprintln(w, "login failed")
		return
	}

	fmt.Fprintln(w, auth.NewToken(1))
}

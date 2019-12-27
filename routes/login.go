package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
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

	user, err := models.FindUser(r.Context(), h.DB, username, models.UserColumns.Username, models.UserColumns.Email, models.UserColumns.PasswordHash, models.UserColumns.Active)

	if err != nil {
		h.Logger.Printf("Could not user : '%s'", username)
		fmt.Fprintf(w, "User '%s' was not found", username)
		return
	}

	err = auth.PasswordHashCompare(user.PasswordHash, password)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Wrong password for '%s'", username)
		return
	}

	fmt.Fprintln(w, h.JWT.NewToken(1))
}

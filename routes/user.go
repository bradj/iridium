package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/volatiletech/sqlboiler/boil"
)

func (h HTTP) userMount() http.Handler {
	r := chi.NewRouter()

	r.Group(publicUser)
	r.Group(protectedUser)

	return r
}

func protectedUser(r chi.Router) {
	r.Use(jwtauth.Verifier(h.JWT.TokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Get("/", h.userGet)
}

func publicUser(r chi.Router) {
	r.Post("/", h.userPost)
}

func (h HTTP) newUser(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/new_user.html")

	if err != nil {
		fmt.Fprintln(w, "new user")
		return
	}

	tmpl.Execute(w, nil)
}

func (h HTTP) userGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "user get")
}

func (h HTTP) userPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password_confirm")

	if password != passwordConfirm {
		fmt.Fprintln(w, "passwords do not match")
		return
	}

	hash, err := auth.GeneratePasswordHash(password)

	if err != nil {
		fmt.Fprintln(w, "error hashing password")
		h.App.Logger.Println(err)
		return
	}

	var user models.User

	user.Username = username
	user.Email = email
	user.PasswordHash = hash

	err = user.Insert(r.Context(), h.DB, boil.Infer())

	if err != nil {
		fmt.Fprintln(w, "error inserting user")
		h.App.Logger.Println(err)
	}
}

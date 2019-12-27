package routes

import (
	"errors"
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

	r.Get("/", Helper.Wrap(h.userGet))
}

func publicUser(r chi.Router) {
	r.Post("/", Helper.Wrap(h.userPost))
}

func (h HTTP) newUser(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("templates/new_user.html")

	if err != nil {
		return err
	}

	tmpl.Execute(w, nil)

	return nil
}

func (h HTTP) userGet(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "user get")

	return nil
}

func (h HTTP) userPost(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password_confirm")

	if password != passwordConfirm {
		return errors.New("passwords do not match")
	}

	hash, err := auth.GeneratePasswordHash(password)

	if err != nil {
		return err
	}

	var user models.User

	user.Username = username
	user.Email = email
	user.PasswordHash = hash

	err = user.Insert(r.Context(), h.DB, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

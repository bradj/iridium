package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
	"github.com/go-chi/chi"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (h HTTP) loginMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", Helper.Wrap(h.loginGet))
	r.Post("/", Helper.Wrap(h.loginPost))

	return r
}

func (h HTTP) loginGet(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("templates/login.html")

	if err != nil {
		return err
	}

	tmpl.Execute(w, nil)
	return nil
}

func (h HTTP) loginPost(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	h.App.Logger.Printf("login request for user '%v'", username)

	user, err := models.Users(qm.Where("username=?", username)).One(r.Context(), h.DB)

	if err != nil {
		return err
	}

	err = auth.PasswordHashCompare(user.PasswordHash, password)

	if err != nil {
		return err
	}

	fmt.Fprintln(w, h.JWT.NewToken(1))
	return nil
}

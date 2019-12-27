package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h HTTP) logoutMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", Helper.Wrap(h.logoutGet))
	r.Post("/", Helper.Wrap(h.logoutPost))

	return r
}

func (h HTTP) logoutGet(w http.ResponseWriter, r *http.Request) error {
	log.Println("logout request")
	fmt.Fprintln(w, "logout")

	return nil
}

func (h HTTP) logoutPost(w http.ResponseWriter, r *http.Request) error {
	log.Println("logout request")
	fmt.Fprintln(w, "logout")

	return nil
}

package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h HTTP) logoutMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.logoutGet)
	r.Post("/", h.logoutPost)

	return r
}

func (h HTTP) logoutGet(w http.ResponseWriter, r *http.Request) {
	log.Println("logout request")
	fmt.Fprintln(w, "logout")
}

func (h HTTP) logoutPost(w http.ResponseWriter, r *http.Request) {
	log.Println("logout request")
	fmt.Fprintln(w, "logout")
}

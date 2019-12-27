package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func (h HTTP) adminMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", Helper.Wrap(h.adminGet))

	// r.Use(AdminOnly)
	// r.Get("/accounts", adminListAccounts)
	return r
}

func (h HTTP) adminGet(w http.ResponseWriter, r *http.Request) error {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))

	return nil
}

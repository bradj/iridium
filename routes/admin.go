package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func (h HTTP) adminMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
	})

	// r.Use(AdminOnly)
	// r.Get("/accounts", adminListAccounts)
	return r
}

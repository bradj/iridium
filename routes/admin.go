package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func (h HTTP) AdminHandler(r chi.Router) {
	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
	})
}

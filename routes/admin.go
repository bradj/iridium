package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
)

func (h HTTP) adminGet(w http.ResponseWriter, r *http.Request) error {
	_, claims, _ := jwtauth.FromContext(r.Context())

	fmt.Fprintf(w, "protected area. hi %v", claims["user_id"])

	return nil
}

package routes

import (
	"fmt"
	"net/http"

	"github.com/bradj/iridium/auth"
)

func (h HTTP) adminGet(w http.ResponseWriter, r *http.Request) error {
	claims := auth.GetClaims(r)

	fmt.Fprintf(w, "protected area. hi %v", claims.Subject)

	return nil
}

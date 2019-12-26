package routes

import (
	"fmt"
	"log"
	"net/http"
)

// LoginHandler handles user logins
func (h HTTP) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Login request")
	fmt.Fprintln(w, "login")
}

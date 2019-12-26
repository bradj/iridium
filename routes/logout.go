package routes

import (
	"fmt"
	"log"
	"net/http"
)

// LogoutHandler handles user logout
func (h HTTP) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout request")
	fmt.Fprintln(w, "logout")
}

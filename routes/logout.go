package routes

import (
	"fmt"
	"log"
	"net/http"
)

// LogoutHandler handles user logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout request")
	fmt.Fprintln(w, "logout")
}

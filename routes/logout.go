package routes

import (
	"fmt"
	"log"
	"net/http"
)

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

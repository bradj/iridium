package main

import (
	"fmt"
	"log"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout request")
	fmt.Fprintln(w, "logout")
}

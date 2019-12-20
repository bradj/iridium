package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Login request")
	fmt.Fprintln(w, "login")
}

package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

// HomeHandler renders the landing page
func (h HTTP) homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprintln(w, "home")
		return
	}

	tmpl.Execute(w, nil)
}

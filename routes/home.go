package routes

import (
	"html/template"
	"net/http"
)

// HomeHandler renders the landing page
func (h HTTP) homeHandler(w http.ResponseWriter, r *http.Request) error {
	h.App.Logger.Printf("home request")

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		return err
	}

	tmpl.Execute(w, nil)
	return nil
}

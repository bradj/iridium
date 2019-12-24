package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// CtxKey is used for context lookups
type CtxKey string

// ConfigKey is the context key for the global config object
const ConfigKey CtxKey = "iridium-config-key"

// DBKey is the context key for the global config object
const DBKey CtxKey = "iridium-db-key"

// HomeHandler renders the landing page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	log.Printf("home db : %#v", r.Context().Value(DBKey))
	log.Printf("home config : %#v", r.Context().Value(ConfigKey))

	if err != nil {
		fmt.Fprintln(w, "home")
		return
	}

	tmpl.Execute(w, nil)
}

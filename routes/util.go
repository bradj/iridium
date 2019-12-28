package routes

import (
	"log"
	"net/http"
)

// ErrorHandler is a type that has dependencies
// for error handling. Like loggers and maybe other
// things like 404 HTML templates or something.
type ErrorHandler struct {
	Logger *log.Logger
}

// Wrap wraps an error-ified http.HandlerFunc and returns a normal http.Handler
func (e ErrorHandler) Wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)

		if err == nil {
			return
		}

		e.Logger.Println("error happened", err)
	}
}

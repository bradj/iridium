package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// ErrorHandler is a type that has dependencies
// for error handling. Like loggers and maybe other
// things like 404 HTML templates or something.
type ErrorHandler struct {
	Logger *log.Logger
}

type ErrResponse struct {
	Error string
}

// Wrap wraps an error-ified http.HandlerFunc and returns a normal http.Handler
func (e ErrorHandler) Wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Logger.Printf("User Agent %s", r.UserAgent())

		err := fn(w, r)

		if err == nil {
			return
		}

		buf, _ := json.Marshal(ErrResponse{Error: err.Error()})

		e.Logger.Printf("error happened %v %T", err, err)

		if errors.Is(err, ErrIncorrectPassword) || errors.Is(err, ErrBadUser) {
			w.WriteHeader(400)
			w.Write(buf)
			return
		}
	}
}

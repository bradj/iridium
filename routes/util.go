package routes

import (
	"context"
	"log"
	"math/rand"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/persistence"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func RandString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func PopulateCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := auth.GetClaims(r)

		user, err := persistence.GetUser(claims.UserId, r.Context(), h.DB)

		if err != nil {
			h.Logger.Print("Could not poulate context with user", err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

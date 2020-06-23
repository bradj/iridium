package routes

import (
	"context"
	"log"
	"math/rand"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
	"github.com/bradj/iridium/persistence"
	"github.com/go-chi/chi"
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

		user, err := persistence.GetUserById(claims.UserId, r.Context(), h.DB)

		if err != nil {
			h.Logger.Print("Error retrieving auth'd user", err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PopulateTargetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		user, err := persistence.GetUserByUsername(username, r.Context(), h.DB)

		if err != nil {
			h.Logger.Print("Error retrieving target user", err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx := context.WithValue(r.Context(), targetUserCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTargetUser(r *http.Request) *models.User {
	return r.Context().Value(targetUserCtxKey).(*models.User)
}

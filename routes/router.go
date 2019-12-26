package routes

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	tokenAuth *jwtauth.JWTAuth
	h         HTTP
)

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

// NewRoutes creates all the routes
func NewRoutes(r chi.Router, a App) {
	h = HTTP{App: a}

	r.Group(protected)
	r.Group(public)
}

// Protected renders all routes requiring auth
func protected(r chi.Router) {
	// Seek, verify and validate JWT tokens
	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Mount("/upload", h.uploadMount())
	r.Mount("/admin", h.adminMount())
}

// Public renders all public routes
func public(r chi.Router) {
	r.Get("/", h.homeHandler)
	r.Get("/login", h.loginHandler)
	r.Get("/logout", h.logoutHandler)
}

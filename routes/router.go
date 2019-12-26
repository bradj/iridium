package routes

import (
	"github.com/bradj/iridium/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	h HTTP
)

// NewRoutes creates all the routes
func NewRoutes(r chi.Router, a App) {
	h = HTTP{App: a}

	r.Group(protected)
	r.Group(public)
}

// Protected renders all routes requiring auth
func protected(r chi.Router) {
	// Seek, verify and validate JWT tokens
	r.Use(jwtauth.Verifier(auth.TokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Mount("/upload", h.uploadMount())
	r.Mount("/admin", h.adminMount())
}

// Public renders all public routes
func public(r chi.Router) {
	r.Get("/", h.homeHandler)

	r.Mount("/login", h.loginMount())
	r.Mount("/logout", h.logoutMount())
}

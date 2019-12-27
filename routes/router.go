package routes

import (
	"github.com/bradj/iridium/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	h HTTP
	// Helper is a utililty variable that provides helpful methods for routing
	Helper ErrorHandler
)

// NewRoutes creates all the routes
func NewRoutes(r chi.Router, a App) {
	h = HTTP{
		App: a,
		JWT: auth.NewJWT(),
	}

	Helper = ErrorHandler{Logger: h.Logger}

	r.Mount("/user", h.userMount())

	r.Group(public)
	r.Group(protected)
}

// Protected renders all routes requiring auth
func protected(r chi.Router) {
	// Seek, verify and validate JWT tokens
	r.Use(jwtauth.Verifier(h.JWT.TokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Mount("/upload", h.uploadMount())
	r.Mount("/admin", h.adminMount())
}

// Public renders all public routes
func public(r chi.Router) {
	r.Get("/", Helper.Wrap(h.homeHandler))
	r.Get("/newuser", Helper.Wrap(h.newUser))

	r.Mount("/login", h.loginMount())
	r.Mount("/logout", h.logoutMount())
}

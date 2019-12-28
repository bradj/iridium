package routes

import (
	"github.com/bradj/iridium/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	h HTTP
	// helper is a utililty variable that provides helpful methods for routing
	helper ErrorHandler
)

// NewRoutes creates all the routes
func NewRoutes(r chi.Router, a App) {
	h = HTTP{
		App: a,
		JWT: auth.NewJWT(),
	}

	helper = ErrorHandler{Logger: h.Logger}

	r.Group(public)
	r.Group(protected)
}

// Protected renders all routes requiring auth
func protected(r chi.Router) {
	// Seek, verify and validate JWT tokens
	r.Use(jwtauth.Verifier(h.JWT.TokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Get("/user", helper.Wrap(h.userGet))

	r.Get("/upload", helper.Wrap(h.uploadGet))
	r.Post("/upload", helper.Wrap(h.uploadPost))

	r.Get("/admin", helper.Wrap(h.adminGet))

	r.Get("/logout", helper.Wrap(h.logoutGet))
	r.Post("/logout", helper.Wrap(h.logoutPost))
}

// Public renders all public routes
func public(r chi.Router) {
	r.Get("/", helper.Wrap(h.homeHandler))

	r.Post("/user", helper.Wrap(h.userPost))  // user creation
	r.Get("/newuser", helper.Wrap(h.newUser)) // user creation

	r.Get("/login", helper.Wrap(h.loginGet))
	r.Post("/login", helper.Wrap(h.loginPost))
}

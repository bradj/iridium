package routes

import (
	"log"
	"os"

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
		App:    a,
		JWT:    auth.NewJWT(),
		Logger: log.New(os.Stdout, "", log.LstdFlags+log.LUTC),
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

	r.Get("/image", helper.Wrap(h.uploadGet))
	r.Post("/image", helper.Wrap(h.uploadPost))

	r.Get("/admin", helper.Wrap(h.adminGet))

	r.Post("/logout", helper.Wrap(h.logoutPost))
}

// Public endpoints
func public(r chi.Router) {
	r.Post("/user", helper.Wrap(h.userPost)) // user creation
	r.Post("/login", helper.Wrap(h.loginPost))
}

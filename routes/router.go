package routes

import (
	"log"
	"os"

	"github.com/bradj/iridium/auth"
	"github.com/go-chi/chi"
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
		Logger: log.New(os.Stdout, "APP-", log.LstdFlags+log.LUTC),
	}

	helper = ErrorHandler{Logger: h.Logger}

	r.Group(public)
	r.Group(protected)
}

// Protected renders all routes requiring auth
func protected(r chi.Router) {
	// Seek, verify and validate JWT header token
	r.Use(auth.Verify, auth.Authenticator)

	r.Get("/user", helper.Wrap(h.userGet))
	r.Get("/user/images", helper.Wrap(h.userGetImages))
	r.Post("/user/images", helper.Wrap(h.userUploadImage))

	r.Get("/admin", helper.Wrap(h.adminGet))

	r.Post("/logout", helper.Wrap(h.logoutPost))
}

// Public endpoints
func public(r chi.Router) {
	r.Post("/user", helper.Wrap(h.userPost)) // user creation
	r.Post("/login", helper.Wrap(h.loginPost))
}

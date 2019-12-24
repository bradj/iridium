package routes

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	tokenAuth *jwtauth.JWTAuth
)

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

// Protected renders all routes requiring auth
func Protected(r chi.Router) {
	// Seek, verify and validate JWT tokens
	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Post("/upload", UploadHandler)
	r.Group(AdminHandler)
}

// Public renders all public routes
func Public(r chi.Router) {
	r.Get("/", HomeHandler)
	r.Get("/login", LoginHandler)
	r.Get("/logout", LogoutHandler)
}

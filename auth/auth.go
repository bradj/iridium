package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var (
	// TokenAuth is the JWT authenticator
	TokenAuth *jwtauth.JWTAuth
)

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("fj98jklsns,nv982nvjkfjdsf903290f3jslk;fj"), nil)
}

// NewToken generates a new JWT
func NewToken(userID int) string {
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user_id": userID})

	return tokenString
}

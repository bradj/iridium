package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type contextKey struct {
	name string
}

type IridiumClaims struct {
	UserId string `json:"userid,omitempty"`
	jwt.StandardClaims
}

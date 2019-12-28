package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

// JWT holds properties for our token auth
type JWT struct {
	TokenAuth *jwtauth.JWTAuth
}

// NewJWT creates a new Auth struct
func NewJWT() JWT {
	return JWT{
		TokenAuth: jwtauth.New("HS256", []byte("fj98jklsns,nv982nvjkfjdsf903290f3jslk;fj"), nil),
	}
}

// NewToken generates a new JWT
func (j JWT) NewToken(userID int) string {
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := j.TokenAuth.Encode(jwt.MapClaims{"user_id": userID})

	return tokenString
}

// GeneratePasswordHash creates a hash from a password
func GeneratePasswordHash(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return hash, nil
}

// PasswordHashCompare compares a password with hash
func PasswordHashCompare(hashedPassword []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

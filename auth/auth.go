package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Context keys
var (
	tokenCtxKey   = &contextKey{"Token"}
	claimCtxKey   = &contextKey{"Claims"}
	authHeaderKey = "IRIDIUM_AUTH"
	secretSignKey = []byte("fj98jklsns,nv982nvjkfjdsf903290f3jslk;fj")
	Logger        = log.New(os.Stdout, "AUTH-", log.LstdFlags+log.LUTC)
)

func tokenFromHeader(r *http.Request) string {
	return r.Header.Get(authHeaderKey)
}

func GetClaims(r *http.Request) *IridiumClaims {
	return r.Context().Value(claimCtxKey).(*IridiumClaims)
}

// Verify the request has a token
func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenString := tokenFromHeader(r)

		if tokenString == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		token, err := parseToken(tokenString)

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx = context.WithValue(ctx, tokenCtxKey, token)
		ctx = context.WithValue(ctx, claimCtxKey, token.Claims.(*IridiumClaims))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &IridiumClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		// TODO : Validate the alg is what I expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return secretSignKey, nil
	})

	if err != nil {
		Logger.Printf("Could not parse token %s", err)
		return nil, err
	}

	return token, nil
}

// NewToken generates a new JWT
func NewToken(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, IridiumClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 86400000, // 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretSignKey)

	if err != nil {
		return ""
	}

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

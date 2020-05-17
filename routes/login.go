package routes

import (
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
)

type loginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	Error string
}

func (h HTTP) loginPost(w http.ResponseWriter, r *http.Request) error {
	var lr loginRequest

	h.bodyDecode(r.Body, &lr)

	h.Logger.Printf("login request for user '%s'", lr.Username)

	user, err := models.Users(
		models.UserWhere.Username.EQ(lr.Username),
	).One(r.Context(), h.DB)

	if err != nil {
		h.Logger.Printf("Error while retrieving login username %v", err)
		return ErrBadUser
	}

	err = auth.PasswordHashCompare(user.PasswordHash, lr.Password)

	if err != nil {
		h.Logger.Printf("Password mismatch %v", err)
		return ErrIncorrectPassword
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "jwt",
		Value:  h.JWT.NewToken(user.ID),
		Domain: "localhost",
	})

	return nil
}

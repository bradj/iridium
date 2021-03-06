package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
)

type loginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	Token string
}

func (h HTTP) loginPost(w http.ResponseWriter, r *http.Request) error {
	var lr loginRequest

	err := h.bodyDecode(r.Body, &lr)

	if err != nil {
		h.Logger.Printf("Could not decode body: %s", err)
		return err
	}

	h.Logger.Printf("login request for user '%s'", lr.Username)

	user, err := models.Users(
		models.UserWhere.Username.EQ(lr.Username),
	).One(r.Context(), h.DB)

	if err != nil {
		return err
	}

	err = auth.PasswordHashCompare(user.PasswordHash, lr.Password)

	if err != nil {
		h.Logger.Printf("Password compare failed %v", err)
		return err
	}

	buf, err := json.Marshal(loginResponse{Token: auth.NewToken(user.ID)})

	if err != nil {
		return err
	}

	w.Write(buf)

	return nil
}

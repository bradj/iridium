package routes

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (h HTTP) userGet(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "user get")

	return nil
}

func (h HTTP) userPost(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password_confirm")

	if password != passwordConfirm {
		return errors.New("passwords do not match")
	}

	hash, err := auth.GeneratePasswordHash(password)

	if err != nil {
		return err
	}

	var user models.User

	user.Username = username
	user.Email = email
	user.PasswordHash = hash

	err = user.Insert(r.Context(), h.DB, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

func (h HTTP) userGetImages(w http.ResponseWriter, r *http.Request) error {
	// user := r.Context().Value("User").(*auth.AuthUser)
	// err := models.Upload(models.UploadWhere.ID.EQ(user.UserId))

	fmt.Fprintln(w, "retrieves user images")

	return nil
}

func (h HTTP) userUploadImage(w http.ResponseWriter, r *http.Request) error {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("upload") // Retrieve the file from form data

	if err != nil {
		h.Logger.Println("Failed to retrieve file from request", err)
		return err
	}

	// Close the file when we finish
	defer file.Close()

	// create hash from file contents
	fileLocation := fmt.Sprintf("uploaded/%v-%v", RandString(128), handler.Filename)

	// Writing the file to disk
	f, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	// Copy the file to the destination path
	_, err = io.Copy(f, file)

	if err != nil {
		return err
	}

	// create upload record
	claims := auth.GetClaims(r)

	var upload models.Upload

	upload.UserID = claims.UserId
	upload.Location = fileLocation
	upload.Type = models.UploadTypeImage

	h.Logger.Printf("Creating upload for user %d", upload.UserID)

	err = upload.Insert(r.Context(), h.DB, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/models"
	"github.com/go-chi/chi"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (h HTTP) userGet(w http.ResponseWriter, r *http.Request) error {
	username := chi.URLParam(r, "username")

	fmt.Fprintln(w, fmt.Sprintf("user get %s", username))

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
	// TODO : Add pagination
	claims := auth.GetClaims(r)
	user := getTargetUser(r)

	if user == nil {
		return errors.New("Could not retrieve target user")
	}

	h.Logger.Printf("Retrieving images for user %s", user.Username)

	uploads, err := models.Uploads(
		qm.OrderBy(models.UploadColumns.CreatedAt),
		qm.InnerJoin("users u on u.id = uploads.user_id"),
		models.UploadWhere.UserID.EQ(claims.UserId)).All(r.Context(), h.DB)

	if err != nil {
		return err
	}

	arr := make([]string, len(uploads))

	for ii, u := range uploads {
		arr[ii] = u.Location
	}

	imageLocations, err := json.Marshal(arr)

	if err != nil {
		h.Logger.Printf("Could not send list of images")
		return err
	}

	w.Write(imageLocations)

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

	h.Logger.Printf("Creating upload for user %s", upload.UserID)

	err = upload.Insert(r.Context(), h.DB, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

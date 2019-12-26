package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func (h HTTP) uploadMount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.uploadGet)
	r.Post("/", h.uploadPost)

	return r
}

func (h HTTP) uploadGet(w http.ResponseWriter, r *http.Request) {
	log.Println("upload request")
	fmt.Fprintln(w, "upload")
}

func (h HTTP) uploadPost(w http.ResponseWriter, r *http.Request) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("upload") // Retrieve the file from form data

	if err != nil {
		panic(err)
	}

	defer file.Close() // Close the file when we finish

	// This is path which we want to store the file
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	// Copy the file to the destination path
	_, err = io.Copy(f, file)

	if err != nil {
		panic(err)
	}
}

package main

import (
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
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
	io.Copy(f, file)
}

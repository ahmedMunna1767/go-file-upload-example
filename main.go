package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// File Api handler
func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Setting max file size 10 MB
	r.ParseMultipartForm(10 << 20)

	// Read the file from form
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	// Close the file at the end
	defer file.Close()

	// Printing File Info
	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
	fmt.Printf("File Size: %+v\n", fileHeader.Size)
	fmt.Printf("MIME Header: %+v\n", fileHeader.Header)

	// Creating the file on the server
	tempFile, err := os.OpenFile(filepath.Join("uploads", fileHeader.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Server Started")
	setupRoutes()
}

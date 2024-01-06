package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var fileName = "theFile"

func beforeEveryRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("http %s request on %s\n", r.Method, r.URL)
}

// returns true if file upload was sucessful
func uploadFile(r *http.Request) bool {
	r.ParseMultipartForm(200 << 20) //200 MB
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return false
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer tempFile.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	return true
}

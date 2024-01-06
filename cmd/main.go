package main

import (
	"fmt"
	"net/http"
)

var port = "5001"

func invalidRouteHandler(w http.ResponseWriter, r *http.Request) {
	beforeEveryRequest(w, r)
	fmt.Printf("Invalid route called %s\n", r.Host)
	return
}

// allows the client to upload a file to use (we save the file)
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	beforeEveryRequest(w, r)
	if r.Method != http.MethodPost {
		return
	}
	success := uploadFile(r)
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "failed to upload the file", http.StatusInternalServerError)
	}
	return
}

// sends the file to the client to download
func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	beforeEveryRequest(w, r)
	http.ServeFile(w, r, fileName)
	return
}

func ping(w http.ResponseWriter, r *http.Request) {
	beforeEveryRequest(w, r)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "pong")
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadFileHandler)
	mux.HandleFunc("/download", downloadFileHandler)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/", invalidRouteHandler)
	fmt.Printf("running on port %s\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}
}

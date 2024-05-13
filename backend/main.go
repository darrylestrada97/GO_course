package main

import (
	"io"
	"net/http"
	"os"
)

func checkAPIKey(r *http.Request) bool {
	// Replace "your-api-key" with your actual API key
	return r.Header.Get("X-API-KEY") == "your-api-key"
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if !checkAPIKey(r) {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	out, err := os.Create("/tmp/uploaded-file")
	if err != nil {
		http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
		return
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
		}
	}(out)

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error writing the file", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("File uploaded successfully"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/upload", handlePost)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

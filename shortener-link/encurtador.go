package main

import (
	"fmt"
	"net/http"
)

var links = make(map[string]string)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/redirect", handleRedirect)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the link shortener!")
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	// Generate a unique key for the original URL
	key := generateKey()

	// Store the original URL and key in the links map
	links[key] = originalURL

	// Return the shortened link to the client
	shortenedLink := fmt.Sprintf("http://localhost:8080/redirect/%s", key)
	fmt.Fprintf(w, shortenedLink)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/redirect/"):]
	originalURL, ok := links[key]
	if !ok {
		http.Error(w, "Invalid key", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}

func generateKey() string {
	// This function would generate a unique key for the original URL
	// A simple implementation would be to use a counter variable
	// and increment it every time a new key is needed
	// In a production system, you may want to use a more robust key generation algorithm
	return "1234"
}

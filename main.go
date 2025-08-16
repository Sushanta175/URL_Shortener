package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
)

var UrlStore = make(map[string]string)
var mu sync.Mutex

func generateShortCode(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b)[:n], nil
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	longUrl := r.URL.Query().Get("url")
	if longUrl == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}

	code, err := generateShortCode(6)
	if err != nil {
		http.Error(w, "failed to generate short code", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	UrlStore[code] = longUrl
	mu.Unlock()

	shortUrl := fmt.Sprintf("http://localhost:8080/%s", code)
	fmt.Fprintf(w, "Shortened URL: %s", shortUrl)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	mu.Lock()
	longUrl, ok := UrlStore[code]
	mu.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longUrl, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", ShortenHandler)
	http.HandleFunc("/", RedirectHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

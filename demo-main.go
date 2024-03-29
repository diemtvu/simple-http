package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n") + "\n"
}

type mainHandler struct{}

func (h *mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %v", r.URL)
	message := formatRequest(r)
	w.Write([]byte(message))
}

type healthzHandler struct{}

func (h *healthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Health check OK")
	w.Write([]byte("OK"))
}

func main() {
	go func() {
		http.ListenAndServe(":8080", &healthzHandler{})
	}()

	http.ListenAndServe(":8000", &mainHandler{})
}

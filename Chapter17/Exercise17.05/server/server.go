package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

// ServeHTTP implements http.Handler.
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Authorization token not recognized"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}
	time.Sleep(10 * time.Second)
	msg := "hello client!"
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}

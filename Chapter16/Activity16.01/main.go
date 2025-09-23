package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageCounter struct {
	Content string
	Heading string
	Counter int
}

func (h *PageCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Counter += 1
	msg := fmt.Sprintf("<h1>%s</h1><p>%s</p><p>Views: %d</p>", h.Heading, h.Content, h.Counter)
	w.Write([]byte(msg))
}

func main() {
	http.Handle("/", &PageCounter{Content: "This is the main page", Heading: "Hello World!"})
	http.Handle("/chapter1", &PageCounter{Content: "This is the first chapter", Heading: "Chapter 1"})
	http.Handle("/chapter2", &PageCounter{Content: "This is the second chapter", Heading: "Chapter 2"})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

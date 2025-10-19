package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `
{"names": ["Electric","Electric","Electric","Boogaloo","Booga-loo","Boogaloo","Boogaloo"]}
	`
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}

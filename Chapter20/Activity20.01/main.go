package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ExampleHandler handles the http requests sent to this webserver
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello Packt")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ExampleHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}

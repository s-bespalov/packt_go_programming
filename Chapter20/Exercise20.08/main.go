package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello Packt")
	if err != nil {
		log.Println("Error in exampleHandler:", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", exampleHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}
type messageData struct {
	Message string `json:"message"`
}

// ServeHTTP implements http.Handler.
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	messageData := messageData{}
	err := jsonDecoder.Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, _ := json.Marshal(messageData)
	log.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}

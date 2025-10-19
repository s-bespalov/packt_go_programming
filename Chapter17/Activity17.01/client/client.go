package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type namesData struct {
	Names []string `json:"names"`
}

func getDatatAndReturnResponse() namesData {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := namesData{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}
	return names
}

func histogram(data []string) map[string]int {
	h := make(map[string]int)
	for _, v := range data {
		h[v] += 1
	}
	return h
}

func main() {
	data := getDatatAndReturnResponse()
	fmt.Println(data.Names)
	h := histogram(data.Names)
	for k, v := range h {
		fmt.Println(k, v)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	r, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data)
}

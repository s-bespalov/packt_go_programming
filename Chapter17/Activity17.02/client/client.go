package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type namesData struct {
	Names []string `json:"names"`
}

type postResults struct {
	Ok bool `json:"ok"`
}

type reqParams struct {
	Name string `json:"name"`
}

func addNameAndParseResponse(name string) (postResults, error) {
	params := reqParams{Name: name}
	results := postResults{}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return results, fmt.Errorf("failed to marshal params: %w", err)
	}
	r, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return results, fmt.Errorf("POST request failew: %w", err)
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			fmt.Println("error when closing request body:", err)
		}
	}()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return results, fmt.Errorf("error when read request body: %w", err)
	}
	err = json.Unmarshal(data, &results)
	if err != nil {
		return results, fmt.Errorf("error when unmarshal response data json: %w", err)
	}
	return results, nil
}

func getDatatAndReturnResponse() namesData {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("Error closing response body:", err)
		}
	}()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := namesData{}
	log.Println(string(data))
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
	rsl, err := addNameAndParseResponse("Aragorn")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Adding name result:", rsl.Ok)
	}
	data := getDatatAndReturnResponse()
	fmt.Println(data.Names)
	h := histogram(data.Names)
	for k, v := range h {
		fmt.Println(k, v)
	}
}

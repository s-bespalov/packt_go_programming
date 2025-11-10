package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	server struct{}
	data   struct {
		Names []string `json:"names,omitempty"`
	}
)

type requestParams struct {
	Name string `json:"name"`
}

type requestResults struct {
	Ok bool `json:"ok"`
}

var storedData data

func init() {
	n := []string{
		"Electric", "Electric", "Electric", "Boogaloo", "Booga-loo", "Boogaloo", "Boogaloo",
	}
	storedData = data{Names: n}
}

func sendResults(ok bool, w http.ResponseWriter) error {
	rsl := requestResults{Ok: ok}
	jsonBytes, err := json.Marshal(rsl)
	if err != nil {
		return fmt.Errorf("error when marshaling results to json: %w", err)
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		return fmt.Errorf("error when write results %w", err)
	}
	return nil
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		defer func() {
			err := r.Body.Close()
			if err != nil {
				log.Println("err when close request body:", err)
			}
		}()
		paramBytes, err := io.ReadAll(r.Body)
		if err != nil {
			err = sendResults(false, w)
			if err != nil {
				log.Println(err)
			}
			log.Fatalln("error when read reques body:", err)
		}
		params := requestParams{}
		err = json.Unmarshal(paramBytes, &params)
		if err != nil {
			err = sendResults(false, w)
			if err != nil {
				log.Println(err)
			}
			log.Fatalln("error when parsing request params to json:", err)
		}
		storedData.Names = append(storedData.Names, params.Name)
		err = sendResults(true, w)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if r.Method == http.MethodGet {
		msg, err := json.Marshal(storedData)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}

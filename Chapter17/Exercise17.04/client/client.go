package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnResponse(filename string) string {
	fileData := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&fileData)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	formFile, err := multipartWriter.CreateFormFile("MyFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	multipartWriter.Close()
	req, err := http.NewRequest("POST", "http://localhost:8080", &fileData)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := postFileAndReturnResponse("./test.txt")
	fmt.Println(data)
}

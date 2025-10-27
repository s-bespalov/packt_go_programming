package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadedFile, uploadedFileHeader, err := r.FormFile("MyFile")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadedFile.Close()
	fileContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0o600)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(fmt.Appendf(nil, "%s Uploaded!", uploadedFileHeader.Filename))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}

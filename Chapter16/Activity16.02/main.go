package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var indexTmpl *template.Template

type Visitor struct {
	Name string
}

func init() {
	var err error
	indexTmpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("error in creating template:", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	v := Visitor{}
	name, ok := val["name"]
	if ok {
		v.Name = strings.Join(name, ",")
	}
	err := indexTmpl.Execute(w, v)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

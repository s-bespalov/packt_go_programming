package main

import (
	"fmt"
	"net/http"
	"strings"
)

func greet(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	name, ok := vl["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing name"))
		return
	}
	fmt.Fprintf(w, "Hello %s", strings.Join(name, ","))
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

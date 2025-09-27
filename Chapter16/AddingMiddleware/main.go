package main

import (
	"log"
	"net/http"
)

func Hello(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello there,"
		w.Write([]byte(msg))
		next.ServeHTTP(w, r)
	}
}

func Function1(w http.ResponseWriter,
	r *http.Request,
) {
	msg := " this is function 1"
	w.Write([]byte(msg))
}
func Function2(w http.ResponseWriter,
	r *http.Request,
) {
	msg := " and now we are in function 2"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc(
		"/hello1", Hello(Function1))
	http.HandleFunc(
		"/hello2", Hello(Function2))
	log.Fatal(http.ListenAndServe(":8085", nil))
}

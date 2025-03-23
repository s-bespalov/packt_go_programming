package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	region   string
}

var locales map[locale]struct{}

func main() {
	locales = make(map[locale]struct{})
	locales[locale{language: "en", region: "CN"}] = struct{}{}
	locales[locale{language: "en", region: "US"}] = struct{}{}
	locales[locale{language: "ru", region: "RU"}] = struct{}{}
	locales[locale{language: "fr", region: "FR"}] = struct{}{}
	locales[locale{language: "fr", region: "CN"}] = struct{}{}

	if len(os.Args) < 2 {
		fmt.Println("No locale in arguments providet")
		os.Exit(1)
	}
	arr := strings.Split(os.Args[1], "_")
	loc := locale{language: arr[0], region: arr[1]}
	_, supported := locales[loc]
	if supported {
		fmt.Println("Locale is supported")
	} else {
		fmt.Println("Locale is not supported:", os.Args[1])
	}
}

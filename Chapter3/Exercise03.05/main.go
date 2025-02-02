package main

import "fmt"

func main() {
	logLevel := "デバシケ"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}

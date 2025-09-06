package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}
	name := args[1]
	greeting := fmt.Sprintf("Hello, %s! Welcome to the commandline.", name)
	fmt.Println(greeting)
}

package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	uuid := uuid.New()
	quote := quote.Glass()
	fmt.Printf("Generated UUID: %s\n", uuid)
	fmt.Printf("Random quote: %s\n", quote)
}

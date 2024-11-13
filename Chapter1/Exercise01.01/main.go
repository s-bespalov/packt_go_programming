package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}

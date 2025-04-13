package main

import (
	"fmt"
)

func decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}

func main() {
	counter := 4
	x := decrement(counter)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

package main

import "fmt"

func main() {
	for i := range 15 {
		num, result := checkNumbers(i)
		fmt.Printf("Results: %d %s\n", num, result)
	}
}

func checkNumbers(i int) (int, string) {
	switch i % 2 {
	case 0:
		return i, "Even"
	default:
		return i, "Odd"
	}
}

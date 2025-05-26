package main

import (
	"fmt"
)

func main() {
	i1 := []int{1, 32, 5, 8, 10, 11}
	f1 := []float64{32.1, 5.1, 8.1, 10.1, 1.1, 11.1}
	min1 := findMinGeneric(i1)
	min2 := findMinGeneric(f1)
	fmt.Printf("input: %v\nMin: %v\n\n", i1, min1)
	fmt.Printf("input: %v\nMin: %v\n\n", f1, min2)
}

func findMinGeneric[Num int | float64](n []Num) Num {
	if len(n) == 0 {
		return -1
	}
	min := n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	return min
}

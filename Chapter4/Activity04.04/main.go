package main

import "fmt"

func removeFomSlice(value string, slice []string) []string {
	var result []string = slice
	for i, v := range slice {
		if v == value {
			result = append(result[:i], result[i+1:]...)
		}
	}
	return result
}

func main() {
	slice := []string{"Good", "Good", "Bad", "Good", "Good"}
	slice = removeFomSlice("Bad", slice)
	fmt.Println(slice)
}

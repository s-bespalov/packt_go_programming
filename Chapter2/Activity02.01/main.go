package main

import "fmt"

func main() {
	popularWord := ""
	popularCount := 0
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	for word, count := range words {
		if count > popularCount {
			popularCount = count
			popularWord = word
		}
	}
	fmt.Println("Most popular word:", popularWord, "\n", "With a count of:", 4)
}

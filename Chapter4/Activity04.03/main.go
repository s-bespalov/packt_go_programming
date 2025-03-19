package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	days = append(days[len(days)-1:], days[:len(days)-1]...)
	fmt.Println(days)
}

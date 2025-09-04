package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Printf("The current time: %s\n", current.Format(time.ANSIC))
	future := current.Add(6*time.Hour + 6*time.Minute + 6*time.Second)
	fmt.Printf("6 hours, 6 minutes and 6 seconds from now the time will be: %s\n", future.Format(time.ANSIC))
}

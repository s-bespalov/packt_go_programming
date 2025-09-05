package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t := time.Now()
	nyTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalln(err)
	}
	laTime, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("The local current time is:", t.Format(time.ANSIC))
	fmt.Println("The time in New York is:", t.In(nyTime).Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is:", t.In(laTime).Format(time.ANSIC))
}

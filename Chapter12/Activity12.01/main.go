package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	day := t.Day()
	month := t.Month()
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()
	seconds := t.Second()
	fmt.Printf("%02d:%02d:%02d %02d/%02d/%04d\n", hour, minute, seconds, day, month, year)
	fmt.Println(t.Format("15:04:05 02/01/2006"))
}

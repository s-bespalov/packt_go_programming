package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Date(2023, 1, 31, 2, 49, 21, 0, time.UTC)
	hour := strconv.Itoa(t.Hour())
	minute := strconv.Itoa(t.Minute())
	seconds := strconv.Itoa(t.Second())
	year := strconv.Itoa(t.Second())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	fmt.Println(hour + ":" + minute + ":" + seconds + " " + year + "/" + month + "/" + day)
}

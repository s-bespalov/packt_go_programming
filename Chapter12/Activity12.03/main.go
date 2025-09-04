package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Elapsed time for time.Sleep(2 sec) is exactly: " + strconv.FormatFloat(duration.Seconds(), 'f', -1, 64))
}

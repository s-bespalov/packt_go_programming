package main

import (
	"fmt"
	"time"
)

var (
	Debug bool = false
	LogLevel string = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
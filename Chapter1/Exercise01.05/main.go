package main

import (
	"fmt"
	"time"
)

func main()  {
	Debug := false
	logLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, logLevel, startUpTime)
}
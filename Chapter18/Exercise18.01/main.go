package main

import (
	"fmt"
	"time"
)

func sum(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}
	return res
}

func main() {
	var s1, s2 int
	go func() {
		s1 = sum(1, 100)
	}()
	s2 = sum(1, 10)
	time.Sleep(time.Second)
	fmt.Println(s1, s2)
}

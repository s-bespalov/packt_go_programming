package main

import "fmt"

func main() {
	visits := 15
	fmt.Println("First visit\t:", visits == 1)
	fmt.Println("Return visit\t:", visits != 1)
	fmt.Println("Silver member\t:", visits >= 10 && visits < 21)
	fmt.Println("Gold member\t:", visits >= 20 && visits <= 30)
	fmt.Println("Platinum member\t:", visits > 30)
}

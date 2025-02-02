package main

import "fmt"

func main() {
	var cackeTax float32 = 0.99 * 0.075
	var milkTax float32 = 2.75 * 0.015
	var butterTax float32 = 0.87 * 0.02
	total := cackeTax + milkTax + butterTax
	fmt.Println("Sales Tax Total: ", total)
}

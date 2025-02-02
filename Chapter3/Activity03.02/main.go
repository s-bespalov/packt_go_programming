package main

import (
	"fmt"
	"log"
)

func main() {
	var name string
	fmt.Println("Inpunt name")
	_, err := fmt.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	var creditScore int
	fmt.Println("Inpunt credit score")
	_, err = fmt.Scan(&creditScore)
	if err != nil {
		log.Fatal(err)
	}
	var income float64
	fmt.Println("Inpunt income")
	_, err = fmt.Scan(&income)
	if err != nil {
		log.Fatal(err)
	}
	var loanAmount float64
	fmt.Println("Input loan Amount")
	_, err = fmt.Scan(&loanAmount)
	if err != nil {
		log.Fatal(err)
	}
	var loanTerm int
	fmt.Println("Inpunt loan term")
	_, err = fmt.Scan(&loanTerm)
	if err != nil {
		log.Fatal(err)
	}
	if loanTerm%12 != 0 {
		log.Fatal("term of the loan should be divisile by 12 month")
	}
	isGood := creditScore > 450
	var rate float64
	if isGood {
		rate = 0.15
	} else {
		rate = 0.2
	}
	var total float64 = (rate * (float64(loanTerm / 12)) * loanAmount) + loanAmount
	var monthly float64 = total / 12
	var approved bool
	if isGood {
		approved = monthly <= income/12*0.2
	} else {
		approved = monthly <= income/12*0.15
	}
	fmt.Println()
	fmt.Println("Application:", name)
	fmt.Println("----------------------")
	fmt.Println("Credit Score:", creditScore)
	fmt.Println("Loan Amount:", loanAmount)
	fmt.Println("Loan Term:", loanTerm)
	fmt.Println("Monthly Payment:", monthly)
	fmt.Println("Rate:", int(rate*100), "%")
	fmt.Println("Total Cost:", total-loanAmount)
	fmt.Println("Approved:", approved)
}

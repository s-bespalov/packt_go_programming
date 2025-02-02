package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	goodCreditScore  int     = 450
	goodRate         float64 = 0.15
	baseRate         float64 = 0.2
	goodMonthlyLimit float64 = 0.2
	baseMonthlyLimit float64 = 0.1
)

func scanInput(prompt string, dest interface{}) {
	fmt.Println(prompt)
	_, err := fmt.Scan(dest)
	if err != nil {
		log.Fatal(err)
	}
}

func calculateCredit(income, loanAmount float64, loanTerm, creditScore int) (rate, monthly, totalCost float64, approved bool, err error) {
	if loanTerm%12 != 0 {
		err = errors.New("term of the loan should be divisible by 12 month")
		return
	}
	isGood := creditScore >= goodCreditScore
	if isGood {
		rate = goodRate
	} else {
		rate = baseRate
	}
	years := loanTerm / 12
	totalCost = (loanAmount * rate) * float64(years)
	monthly = (loanAmount + totalCost) / float64(loanTerm)
	if isGood {
		approved = (goodMonthlyLimit * income) >= monthly
	} else {
		approved = (baseMonthlyLimit * income) >= monthly
	}
	return
}

func main() {
	var name string
	var creditScore int
	var income float64
	var loanAmount float64
	var loanTerm int
	scanInput("Input name", &name)
	scanInput("Input credit score", &creditScore)
	scanInput("Input monthly income", &income)
	scanInput("loan Amount", &loanAmount)
	scanInput("Input loan term", &loanTerm)
	rate, monthly, totalCost, approved, err := calculateCredit(income, loanAmount, loanTerm, creditScore)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nApplication:", name)
	fmt.Println("----------------------")
	fmt.Println("Credit Score:", creditScore)
	fmt.Println("income:", income)
	fmt.Println("Loan Amount:", loanAmount)
	fmt.Println("Loan Term:", loanTerm)
	fmt.Println("Monthly Payment:", monthly)
	fmt.Println("Rate:", int(rate*100), "%")
	fmt.Println("Total Cost:", totalCost)
	fmt.Println("Approved:", approved)
}

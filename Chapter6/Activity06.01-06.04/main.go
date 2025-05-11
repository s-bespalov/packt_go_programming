package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type DirectDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func main() {
	deposit := DirectDeposit{"", "Abe", "XYZ Inc", 17, 1809}
	err := deposit.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = deposit.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	deposit.report()
}

func (d *DirectDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d *DirectDeposit) validateLastName() error {
	if len(d.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}

func (d *DirectDeposit) report() {
	fmt.Printf("%s\nLast name: %s\nFirst name: %s\nBank name: %s\nRouting number: %d\nAccount Number:%d\n",
		strings.Repeat("*", 80), d.lastName, d.firstName, d.bankName, d.routingNumber, d.accountNumber)
}

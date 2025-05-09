package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("Invalid hourly rate")
	ErrHoursWorked = errors.New("Invalid hours worked per week")
)

func main() {
	pay := payDay(100, 25)
	fmt.Println(pay)
	pay = payDay(100, 200)
	fmt.Println(pay)
	pay = payDay(60, 25)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours worked: %d\nerr: %v\n\n", hoursWorked, r)
			}
		}
		fmt.Printf("Payment was calculated based on:\nhours worked:%d\nhourly rate: %d\n", hoursWorked, hourlyRate)
	}()
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}
	if hoursWorked > 40 {
		return (hoursWorked + (hoursWorked - 40)) * hourlyRate
	}
	return hoursWorked * hourlyRate
}

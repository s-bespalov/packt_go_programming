package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func main() {
	pay := payDay(81, 50)
	fmt.Println(pay)
}

func payDay(HoursWorked, HourlyRate int) int {
	defer func() {
		fmt.Printf("Hours worked: %d, Hourly rate: %d\n", HoursWorked, HourlyRate)
	}()
	if HourlyRate < 10 || HourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if HoursWorked < 0 || HoursWorked > 80 {
		panic(ErrHoursWorked)
	}
	if HoursWorked > 40 {
		return (HoursWorked + (HoursWorked - 40)) * HourlyRate
	}
	return HoursWorked * HourlyRate
}

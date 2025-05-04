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
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}

func payDay(HoursWorked, HourlyRate int) (int, error) {
	if HourlyRate < 10 || HourlyRate > 75 {
		return 0, ErrHourlyRate
	}
	if HoursWorked < 0 || HoursWorked > 80 {
		return 0, ErrHoursWorked
	}
	if HoursWorked > 40 {
		return (HoursWorked + (HoursWorked - 40)) * HourlyRate, nil
	}
	return HoursWorked * HourlyRate, nil
}

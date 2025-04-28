package main

import (
	"fmt"
)

type Weekday int

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (d Weekday) String() string {
	switch d {
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Sunday:
		return "Sunday"
	default:
		return "Unknown"
	}
}

type Employee struct {
	Id        int
	Firstname string
	Lastname  string
}

type Devloper struct {
	Individual     Employee
	HourlyRate     int
	nonLoggedHours int
	WorkWeek       [7]int
}

func (d *Devloper) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Devloper) HoursWorked() int {
	sum := 0
	for _, v := range d.WorkWeek {
		sum += v
	}
	return sum
}

func nonLoggedHours() func(int) int {
	x := 0
	return func(y int) int {
		x += y
		return x
	}
}

func (d *Devloper) PayDetails() {
	for i, v := range &d.WorkWeek {
		if v > 0 {
			fmt.Printf("%s hours: %d\n", Weekday(i), v)
		}
	}
}

func (d *Devloper) PayDay() (int, bool) {
	hours := d.HoursWorked()
	if hours <= 40 {
		return hours * d.HourlyRate, false
	} else {
		return (40 * d.HourlyRate) + ((hours - 40) * d.HourlyRate * 2), true
	}
}

func main() {
	devloper1 := Devloper{HourlyRate: 10}
	nonLoggedHours := nonLoggedHours()
	fmt.Println("Tracking hours working thus far today:", nonLoggedHours(2))
	fmt.Println("Tracking hours working thus far today:", nonLoggedHours(3))
	fmt.Println("Tracking hours working thus far today:", nonLoggedHours(5))
	devloper1.LogHours(Monday, 8)
	devloper1.LogHours(Tuesday, 10)
	devloper1.LogHours(Wednesday, 10)
	devloper1.LogHours(Thursday, 10)
	devloper1.LogHours(Friday, 6)
	devloper1.LogHours(Saturday, 8)
	devloper1.PayDetails()
	fmt.Println("Hoours worked this week:", devloper1.HoursWorked())
	salary, overtime := devloper1.PayDay()
	fmt.Printf("Pay for the week: $%d\n", salary)
	fmt.Println("is this overtime pay:", overtime)
}

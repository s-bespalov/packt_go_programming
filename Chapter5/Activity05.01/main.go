package main

import "fmt"

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

type Employee struct {
	Id        int
	Firstname string
	Lastname  string
}

type Devloper struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
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

func main() {
	devloper1 := Devloper{}
	devloper1.LogHours(Monday, 8)
	devloper1.LogHours(Tuesday, 10)
	fmt.Printf("Hours worked on Monday: %d\n", devloper1.WorkWeek[Monday])
	fmt.Printf("Hours worked on Tuesday: %d\n", devloper1.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week: %d\n", devloper1.HoursWorked())
}

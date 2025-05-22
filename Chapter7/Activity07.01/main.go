package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        int
	HoursWorkedInYear int
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         int
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	return d.Individual.FirstName, float64(d.HourlyRate) * float64(d.HoursWorkedInYear)
}

func (m Manager) Pay() (string, float64) {
	return m.Individual.FirstName, float64(m.Salary) + (float64(m.CommissionRate) * float64(m.Salary))
}

func payDetails(p Payer) {
	name, payment := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", name, payment)
}

func (d Developer) ReviewRating() error {
	var rating, count float64
	for _, v := range d.Review {
		switch r := v.(type) {
		case int:
			rating += float64(r)
			count += 1
		case string:
			switch r {
			case "Excellent":
				rating += 5
				count += 1
			case "Good":
				rating += 4
				count += 1
			case "Fair":
				rating += 3
				count += 1
			case "Poor":
				rating += 2
				count += 1
			case "Unsatisfactory":
				rating += 1
				count += 1
			}
		default:
			return errors.New("uncnown type")
		}
	}
	rating = rating / count
	fmt.Printf("%s got a reviw rating of %.2f\n", d.Individual.FirstName, rating)
	return nil
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}

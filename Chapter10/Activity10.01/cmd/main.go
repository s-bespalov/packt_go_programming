package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/s-bespalov/packt_go_programming/Chapter10/Activity10.01/payroll"
)

var employeeReview = make(map[string]interface{})
var d payroll.Developer
var m payroll.Manager

func init() {
	fmt.Println("Welcome to the Employee Pay and Performance Review")
	fmt.Println(strings.Repeat("+", 50))
}

func init() {
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d = payroll.Developer{Individual: payroll.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m = payroll.Manager{Individual: payroll.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
}

func main() {
	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payroll.PayDetails(d)
	payroll.PayDetails(m)
}

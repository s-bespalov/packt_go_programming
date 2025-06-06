package payroll

import (
	"errors"
	"fmt"
)

type Developer struct {
	Individual        Employee
	HourlyRate        int
	HoursWorkedInYear int
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	return d.Individual.FirstName, float64(d.HourlyRate) * float64(d.HoursWorkedInYear)
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

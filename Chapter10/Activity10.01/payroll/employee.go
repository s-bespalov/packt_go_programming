package payroll

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Payer interface {
	Pay() (string, float64)
}

func PayDetails(p Payer) {
	name, payment := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", name, payment)
}

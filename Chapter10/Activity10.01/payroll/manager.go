package payroll

type Manager struct {
	Individual     Employee
	Salary         int
	CommissionRate float64
}

func (m Manager) Pay() (string, float64) {
	return m.Individual.FirstName, float64(m.Salary) + (float64(m.CommissionRate) * float64(m.Salary))
}

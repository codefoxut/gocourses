package factories

import (
	"fmt"
)

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{Name: name, Position: position, AnnualIncome: annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(postion string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{postion, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

// prototype

const (
	Developer = iota
	Manager
)

// functional
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 10000}
	case Manager:
		return &Employee{"", "Manager", 20000}
	default:
		panic("unsupported role.")
	}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	fmt.Println(*developer)

	manager := managerFactory("Jane")
	fmt.Println(*manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Joshua")
	fmt.Println(*boss)

	boss.AnnualIncome = 100
	fmt.Println(*boss)

	boss2 := bossFactory.Create("Rachel")
	fmt.Println(*boss2)

	// prototype
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(*m)

}

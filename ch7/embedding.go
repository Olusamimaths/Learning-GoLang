package main

import "fmt"

type Employee struct {
	Name string
	ID string
}
func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee	// note no field name here => embedded field
	Reports []Employee
}

func main() {
	m := Manager {
		Employee: Employee{
			Name: "Sam",
			ID: "1234",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Description())
}
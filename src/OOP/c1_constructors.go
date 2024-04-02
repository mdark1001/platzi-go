/**
* This programa tries to demostrate how to use constructors and methods in GO
**/

package main

import "fmt"

type Employee struct {
	id   string
	name string
	job  string
}

func (employee *Employee) String() string {
	return fmt.Sprintf("<Employee (%s, %s, %s)>", employee.id, employee.name, employee.job)
}
func NewEmployee(id, name, job string) *Employee {
	// this kind of functions returns an pointer to a new instance of Employee
	return &Employee{id: id, name: name, job: job}
}

func main() {
	// Constructor using default values
	var employee1 Employee = Employee{}
	fmt.Println(employee1)

	// Constructor passing values

	var employee2 Employee = Employee{id: "123", name: "John", job: "software engineer"}
	fmt.Println(employee2)

	//  constructor using a funcition with values
	var employee3 Employee = *NewEmployee(
		"3412", "Samantha", "Marketing sales manager",
	)
	fmt.Println(employee3)

}

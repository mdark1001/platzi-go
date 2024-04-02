// implementation of interfaces in Go

package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FulltimeEmployee struct {
	// this struct implements the anonimus fields below
	Person
	Employee
}
type TemporalEmployee struct {
	Person
	Employee
	taxRate float64
}

// creating the interface
type PrintInfo interface {
	getMessage() string
}

func (fEmployee *FulltimeEmployee) getMessage() string {
	return fmt.Sprintf("Employee Full time: %d\n", fEmployee.id)
}
func (tEmployee *TemporalEmployee) getMessage() string {
	return fmt.Sprintf("Employee temporal: %d\n", tEmployee.id)
}

// this function implements polimorphisms for PrintInfo's interface subclasses.
func printAllMessages(employeeList []PrintInfo) {
	for _, employee := range employeeList {
		fmt.Print(employee.getMessage())
	}
}

func main() {

	var emp1 = FulltimeEmployee{}
	emp1.id = 100
	emp2 := TemporalEmployee{}
	emp2.id = 123
	var employeeList = []PrintInfo{&emp1, &emp2}

	printAllMessages(employeeList)
}

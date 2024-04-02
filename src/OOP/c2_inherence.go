// This programa is really interested in because
// in go prefers the composition over the inherece

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

func main() {
	// creating a new fulltime employee Passing the Person and Employee
	var ftEmployee1 FulltimeEmployee = FulltimeEmployee{
		Person{
			name: "John", age: 32,
		},
		Employee{
			id: 12,
		},
	}
	fmt.Println(ftEmployee1)

	var ftEmployee2 FulltimeEmployee = FulltimeEmployee{}

	ftEmployee2.name = "John"
	ftEmployee2.age = 35
	ftEmployee2.id = 145

	fmt.Println(ftEmployee2)

}

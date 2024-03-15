/*
@author: mdark1001
@date: 04/03/2024
@description: Creating Strcuts in go
*/
package main

import "fmt"

type Student struct {
	student_id string
	name       string
	grade      float32
}
type Kardex struct {
	calculateGrates float32
}

func (s *Student) calculateGrates() float32 {
	return s.grade * 1.5
}

func main() {
	var students []Student
	students = append(students, Student{
		student_id: "12145-1213",
		name:       "Michael Smith",
		grade:      94.0,
	})
	students = append(students, Student{
		student_id: "67534-1413",
		name:       "John Latt",
		grade:      92.0,
	})

	fmt.Println(students[0])
	fmt.Println(students)
	var otherStudent Student

	otherStudent.student_id = "45753-1212"

	fmt.Println(otherStudent)
}

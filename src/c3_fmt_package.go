/*
@author: mdark1001
@date: 26/02/2024
@description: Using fmt package, more than printLn-
*/
package main

import "fmt"

func main() {

	var helloMessage string = "Hello"
	var worldMessage string = "world!"

	fmt.Println(helloMessage, worldMessage)
	// formatin string output
	fmt.Printf("%s everybody in the %s \n", helloMessage, worldMessage)
	// printing numeric values
	var followers int = 100
	var percentage float32 = 78.50

	fmt.Printf("I have %d followers, I increased more than %f% \n", followers, percentage)
	// if you don't know the type
	fmt.Printf("%v of %v\n", worldMessage, percentage)
	// build a string representation
	var messageRepresentation string = fmt.Sprintf("%s%s has %d followers", helloMessage, worldMessage, followers)

	fmt.Println(messageRepresentation)
	// printing the variable's type
	fmt.Printf("followers is: %T\n", followers)
	fmt.Printf("Percentage is: %T\n", percentage)
}

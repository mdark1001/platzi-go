/*
@author: mdark1001
@date: 11/03/2024
@description: Iterates over a slice of strings using the keyword range
*/
package main

import "fmt"

func main() {

	// Slices
	var slice = []string{"A", "B", "C", "D"}
	fmt.Println(slice)

	for index, value := range slice {
		fmt.Println("Index: ", index, " with value", value)
	}

}

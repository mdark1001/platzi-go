/*
@author: mdark1001
@date: 04/03/2024
@description: Creating for loops in go
*/
package main

import "fmt"

func main() {
	// for conditional
	for i := 0; i <= 10; i++ {
		fmt.Println("Value: ", i)
	}
	// for while
	var counter int = 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}

}

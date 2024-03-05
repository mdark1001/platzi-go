/*
@author: mdark1001
@date: 26/02/2024
@description: Making functions and anomimous functions
*/
package main

import "fmt"

/**
* Sum: Sum of two integers
**/
func sum(x, y int) int {
	return x + y
}
func times(number1, number2 int) int {
	return number1 * number2
}
func div(n, m float64) float64 {
	return float64(n / m)
}

/*
* This is a function that return multiple values.
 */
func multipleReturns(a int) (b int, c string) {
	return a * 2, "The result is..."
}

func main() {
	fmt.Println(sum(13, 19))
	fmt.Println(times(13, 19))
	fmt.Println(div(100.3, 2.5))
	var resInt int
	var resString string
	resInt, resString = multipleReturns(10)
	fmt.Printf("%d %s\n", resInt, resString)
}

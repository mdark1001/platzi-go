/*
@author: mdark1001
@date: 04/03/2024
@description: Creating arrays in go
*/
package main

import "fmt"

func main() {

	var size int = 10
	var mltTable [10]int
	for i := 0; i < size; i++ {
		mltTable[i] = (i + 1) * 3
	}

	fmt.Println("Results")
	for i := 0; i < size; i++ {
		fmt.Println(i+1, "x", 3, "=", mltTable[i])
	}
	// funciones  de los arrays
	fmt.Println(mltTable, len(mltTable), cap(mltTable))

	// Slices
	var slice = []int{0, 1, 2, 3}
	fmt.Println(slice)
	// Agregar elementos al slice
	slice = append(slice, 4)
	fmt.Println(slice)
	// slicing similar a python
	fmt.Println(slice[0:2])
	fmt.Println(slice[3:])
	// descomprimir un slice y agregar un slice a otro
	var newSlice = []int{5, 6, 7}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}

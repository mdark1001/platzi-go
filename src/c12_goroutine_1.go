/*
@author: mdark1001
@date: 04/03/2024
@description: Creating methods for structs  in go
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arrayNumbers = []int{}
	for i := 0; i < 10; i++ {
		arrayNumbers = append(arrayNumbers, (rand.Intn(1000-1) + 1))
	}
	resultados := make(chan int)
	for _, number := range arrayNumbers {
		// fmt.Println(number)
		go func(n int) {
			// fmt.Println(n)
			resultados <- n * n // ingresar datos al canal
		}(number)
	}

	fmt.Println(arrayNumbers)
	defer close(resultados) // clos a channel
	for range arrayNumbers {
		fmt.Println(<-resultados) // take a message of the channel
	}
	fmt.Printf("2, %T", 2)
	// creating a map of arrays' strings
	var g = make(map[int][]string)
	// adding a new string in the map
	g[2020] = append(g[2020], "23")
}

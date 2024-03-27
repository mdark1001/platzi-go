/*
* This pŕogram is an example of how we  can use go routines and
* communicate them using channels.
* This program will build an arrauy 2-dimensional and
* find the minum value of each arrays
 */

package main

import (
	"fmt"
	"math/rand"
)

const MATRIX_SIZE int = 10

func initMatixwithRandomNumber(matrix *[MATRIX_SIZE][MATRIX_SIZE]int) {
	for row, n := range matrix {
		for col, _ := range n {
			matrix[row][col] = rand.Intn(100)
		}
	}
}

func printMKatrx(matrix [MATRIX_SIZE][MATRIX_SIZE]int) {
	for _, n := range matrix {
		fmt.Println(n)
	}
}

func searchMinimunValue(rowIndex int, listValues [MATRIX_SIZE]int, channel, channelIndex chan<- int) {
	var min int = listValues[0] // select the first value
	for _, v := range listValues {
		if v < min {
			min = v
		}
	}
	channel <- min
	channelIndex <- rowIndex
}

func main() {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]int
	// creating tow channels, one of them to receive the min value
	//  and the oher to receive the max value the index of the row
	var channelResults chan int = make(chan int, 2)
	var channelIndex chan int = make(chan int, 2)

	// inicialize the matrix.
	initMatixwithRandomNumber(&matrix)
	// Printing the matrix
	printMKatrx(matrix)

	// create new routines to looking for the minimun value of each row in the matrix
	for row, list := range matrix {
		go searchMinimunValue(row, list, channelResults, channelIndex)
	}

	fmt.Println("\nThe results of the search are:\n")
	// printhe the results
	for i := 0; i < MATRIX_SIZE; i++ {
		fmt.Printf("Index %d - value [ %d ]\n", <-channelIndex, <-channelResults)
	}

}

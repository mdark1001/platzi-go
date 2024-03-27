package main

import "fmt"

func main() {
	// create a channel nbufferd
	var channel chan int = make(chan int,2)

	// sned information throught the channel
	channel <- 1 // we send one

	fmt.Println(<-channel)

}

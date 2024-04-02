/**
* This programa tries to demostrate how use the interfaces in Go
**/
package main

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	run() (bool, error)
	isRunning() bool
}

type Car struct {
	brand      string
	is_running bool
}

func (c *Car) run() (bool, error) {
	if c.is_running == true {
		return false, errors.New("Car is already running")
	} else {
		fmt.Printf("The car now is running")
	}
	return c.is_running, nil

}
func NewCar(brand string) *Car {
	return &Car{
		brand:      brand,
		is_running: false,
	}
}

func main() {
	var fordFiesta *Car = NewCar("Ford")
	fordFiesta.run()
	_, err := fordFiesta.run()
	if err != nil {
		fmt.Println(err)
	}
}

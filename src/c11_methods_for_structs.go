/*
@author: mdark1001
@date: 04/03/2024
@description: Creating methods for structs  in go
*/
package main

import "fmt"

type Car struct {
	brand      string
	year       int
	passangers int
	is_full    bool
}

/*
* Adding string representation for Car struct.
 */
func (c Car) String() string {
	return fmt.Sprintf("<%s (%d)>", c.brand, c.year)
}
func (c *Car) setFull() {
	c.is_full = true
}

func (c Car) isEmpty() bool {
	return !c.is_full
}

func main() {
	var newCar Car = Car{brand: "Ford", year: 2024, passangers: 4}
	fmt.Println(newCar)
	fmt.Println("Is empty", newCar.isEmpty())
	newCar.setFull()
	fmt.Println("Is empty", newCar.isEmpty())
}

/*
@author: mdark1001
@date: 04/03/2024
@description: Creating maps in go
*/
package main

import "fmt"

func main() {
	// Createing a map, witch is has keys' strings and values integers
	var newMap = make(map[string]int)

	newMap["Charly"] = 30
	newMap["Mike"] = 31
	newMap["Jules"] = 29

	fmt.Println(newMap)
	// looping over the keys and values
	for key, value := range newMap {
		fmt.Println("Key: ", key, " with value {", value, "}")
	}
	// checing if a key is setted on the map
	var keyLookingFor string = "randomKey"
	_, ok := newMap[keyLookingFor]
	if !ok {
		fmt.Println("Key is not setted on the map ", keyLookingFor)
	}

}

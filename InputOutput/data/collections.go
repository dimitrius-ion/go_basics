package data

import "fmt"

var Countries [10] string

func init() {
	fmt.Println("Collections")
	fmt.Println("------------")
	// Arrays
	Prices := [2]int {100, 50} 
	// Slices
	names := []string {"Max", "Dan"}
	names = append(names, "Sam")
	fmt.Println("length of array: ", len(names))

	//Maps
	wellKnownPorts := map[string]int {"http": 80, "https": 443}

	Countries[0] = "Argentina"
	Countries[8] = "USA"
	fmt.Println("Collections: ", Countries, Prices, names, wellKnownPorts)
}
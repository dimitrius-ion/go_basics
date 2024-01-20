package main

import "github.com/dimitrius-ion/go_basics/data"

// can use multiple init functions 
func init() {
	println("Init")
}
func init() {
	println("Init2")
}


func calcTax(price float64) (stateTax float64, cityTax float64) {
	return price * .09, price  * 0.05
}


func main() {
	_, cityTax := calcTax(100)
	println(cityTax)

	data.PrintConstant()
	println("Hello World")
}


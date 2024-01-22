package main

import (
	"fmt"
	"github.com/dimitrius-ion/go_basics/fem/data"
)

func main() {
	max := data.Instructor{Id: 3}
	max.FirstName = "Max"
	fmt.Println(max.Print())

	// using factory function
	kyle := data.NewInstructor("Kyle", "Simpson")
	fmt.Println(kyle.Print())
	
	goCourse := data.Course {Id: 2, Name: "Go Fundamentals", Instructor: kyle}
	
	fmt.Printf("%v\n", goCourse)
}

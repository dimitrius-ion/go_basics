package main

import (
	"fmt"
	"github.com/dimitrius-ion/go_basics/fem/data"
	"time"
)

func main() {

	// Instructors 
	max := data.Instructor{Id: 3}
	max.FirstName = "Max"
	fmt.Println(max.Print())

	// using factory function
	kyle := data.NewInstructor("Kyle", "Simpson")
	fmt.Println(kyle.Print())
	

	// Courses
	goCourse := data.Course {Id: 2, Name: "Go Fundamentals", Instructor: kyle}
	fmt.Printf("%v\n", goCourse)


	//Workshops extends Course
	swiftWs := data.Workshop{Course: data.Course {Name: "Swift for iOS", Instructor: max}, Date: time.Now()}
	fmt.Println(swiftWs)

	//Workshops Embedings Course
	swiftWs_emb := data.Workshop_Embedding {}
	swiftWs_emb.Name = "Swift for iOS"

	swiftWs_emb_2 := data.NewWorkshop_Embedding("Swift for iOS", max, "Swift for iOS")
	//methods gets inherited
	fmt.Println(swiftWs_emb, swiftWs_emb_2)


	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWs_emb

	for i, course := range courses {
		fmt.Println(i, course)
	}
}



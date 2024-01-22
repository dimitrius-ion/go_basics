package main

import "fmt"

func birthday(pointerAge *int) {
	fmt.Println("pointerAge: ",*pointerAge, pointerAge)
	if (*pointerAge > 100){
		defer fmt.Println("Defer in birthday")
		// exit program with error but will honnor defer
		panic("Invalid Age")
	}
	*pointerAge ++
}


func init() {
	age :=44
	birthday(&age)
	fmt.Printf("age: %v\n",age)
 
}
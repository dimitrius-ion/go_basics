package data

import "fmt"

func birthday(pointerAge *int) {
	fmt.Println("pointerAge: ",*pointerAge, pointerAge)
	*pointerAge ++
	
}


func init() {
	age :=44
	birthday(&age)
	fmt.Printf("age: %v\n",age)
 
}
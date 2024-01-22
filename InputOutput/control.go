package main

import "fmt"
// func ifStatement () {
// 	// user,err will be local to if else block
// 	if user, err := getUser(123); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(user)
// 	}
// 	fmt.Println('user', user)
// }

func switchStatmetn () {
	day := "Saturday"
	switch day {
		case "Monday":
			fmt. Println("It's Monday!")
		case "Saturday":
			// will fallthrough to Sunday
			fallthrough 
		case "Sunday":
			fmt. Println("It's Weekend")
		default: 
			fmt. Println("It's another working day")
	}
}

func forStatment () {

	names := []string {"Max", "Dan"}
	
	for index, value := range names {
		fmt.Println(index, value)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for i in range 5 {
	// 	fmt.Println(i)
	// }

	// infinite loop
	for {
		fmt.Println("Infinite loop")
	}
	count:= 20
	// while loop
	for  count < 10 {
		fmt.Println(count)
		count++
	}
}

package main
import "fmt"

func main() {
	var operator string
	var number1, number2 int

	fmt.Println("Welcome to Calculator")
	fmt.Println("---------------------")
	fmt.Println("Enter first number")
	fmt.Scanln(&number1)
	fmt.Println("Enter second number")
	fmt.Scanln(&number2)
	fmt.Println("Enter operator")
	fmt.Scanln(&operator)

	switch operator {
		case "+":
			fmt.Println(number1 + number2)
		case "-":
			fmt.Println(number1 - number2)
		case "*":
			fmt.Println(number1 * number2)
		case "/":
			fmt.Println(number1 / number2)
		default:

	}
	

}
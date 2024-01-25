package main
import (
	"fmt"
	"time"
)

func printMessage () {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
		time.Sleep(time.Second * 2)
	}
}

// main runs in a goroutine
func main() {

	// go keyword is used to create a goroutine
	go printMessage()
	printMessage()
			
	
}
package main
import (
	"fmt"
	"time"
)

func printMessage (chanel chan string) {
	
	for i := 0; i < 2; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Second * 2)
	}
	chanel <- "Done"
}

// main runs in a goroutine
func main() {

	// create a channel that are used to communicate between goroutines
	var chanel chan string
	// chanel := make(chan string)
	// buffer_channel := make(chan string, 2)


	// go keyword is used to create a goroutine
	go printMessage(chanel)
	response := <- chanel

	// close the channel to avoid deadlock
	close(chanel)
	
	fmt.Println(response)
	time.Sleep(time.Second * 10)
			
	
}
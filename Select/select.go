package main

import (
	"fmt"
	"time"
)

func main() {
	number := make(chan int)
	message := make(chan string)

	go channelNumber(number)
	go channelMessage(message)

	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data: ", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data: ", secondChannel)

	default:
		fmt.Println("Wait!! Channels are not ready for execution")
	}

}

func channelNumber(number chan int) {
	number <- 15
	time.Sleep(2 * time.Second)

}

func channelMessage(message chan string) {
	time.Sleep(2 * time.Second)

	message <- "Learning to select"
}

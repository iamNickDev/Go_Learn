package main

import "fmt"

func main() {

	// create channel
	number := make(chan int)
	message := make(chan string)
	ch := make(chan string)
	ch1 := make(chan string)

	// function call with goroutine
	go channelData(number, message)
	go sendData(ch)
	go receiveData(ch1)

	// retrieve channel data
	fmt.Println("Channel Data:", <-number)
	fmt.Println("Channel Data:", <-message)
	fmt.Println(<-ch)

	ch <- "Data Received. Receive Operation Succeful"

}

func channelData(number chan int, message chan string) {

	// send data into channel
	number <- 15
	message <- "Learning Go channel"

}

func sendData(ch chan string) {
	ch <- "Received. Send Operation Successful"
	fmt.Println("No receiver! Send Operation Blocked")

}
func receiveData(ch1 chan string) {

	// receive data from the channel
	fmt.Println(<-ch1)

}

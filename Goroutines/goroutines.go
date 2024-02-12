package main

import (
	"fmt"
	"time"
)

func display(message string) {
	fmt.Println(message)
	// for {
	// 	time.Sleep(time.Second * 1)
	// }
}

func main() {
	// go display("Process 1")

	display("Process 2")
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")
	time.Sleep(time.Second * 1)

}

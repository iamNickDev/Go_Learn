package main

import "fmt"

func main() {

	// array of numbers
	numbers := [5]int{21, 22, 23, 24, 25}

	// Use range to iterate over the elements array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}
}

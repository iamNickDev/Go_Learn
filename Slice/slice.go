package main

import "fmt"

func main() {
	// Create slice in Golang
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Create slice from an array in Golang
	arrayInteger := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	// create slice from an array
	sliceNumbers := arrayInteger[4:7]

	fmt.Println(sliceNumbers)

	// Adds Element to a Slice
	primeNumbers := []int{2, 3}
	primeNumbers = append(primeNumbers, 5, 7)
	fmt.Println("Prime Numbers", primeNumbers)

	// copy(primeNumbers, numbers)

	// Copy Golang Slice
	newSlice := []int{2, 3, 5, 7}
	copySlice := []int{1, 2, 3, 7}
	copy(newSlice, copySlice)
	fmt.Println("Slice Numbers", newSlice)

	// find the length of the slice
	length := len(numbers)
	fmt.Println("Length", length)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Loop", numbers[i])
	}
}

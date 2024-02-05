package main

import "fmt"

// outside the function
var sum int

// Go Local Variables
func addNumbers() {
	// var sum int
	sum = 10 + 2
}

func main() {
	addNumbers()
	// cannot access sum out of its local scope
	// fmt.Println("Sum is: ", sum)
	fmt.Println("Plus is", sum)
}

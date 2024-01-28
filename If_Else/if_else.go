package main

import "fmt"

func main() {
	number1 := 15
	number2 := 20

	// Simple if statement in Golang
	if number1 > 0 {
		fmt.Printf("%d is a positive number\n", number1)
	}

	// if...else statement in Golang
	if number1 < 0 {
		fmt.Printf("%d is a positive number\n", number1)
	} else {
		fmt.Printf("%d is a negative number\n", number1)
	}

	// if...else if ladder statement in Golang
	if number1 == number2 {
		fmt.Printf("Result: %d == %d", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}

	// Nested if statement in Golang
	if number1 >= number2 {
		// inner if statement
		if number1 == number2 {
			fmt.Printf("Result: %d == %d", number1, number2)
		} else {
			fmt.Printf("Result: %d > %d", number1, number2)
		}
	} else {
		fmt.Printf("\nResult: %d < %d", number1, number2)
	}
}

package main

import "fmt"

func division(num1, num2 int) {

	defer handlePanic()
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2

		fmt.Println("Result: ", result)
	}
}

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}
}

func main() {
	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")

	defer fmt.Println("four")
	defer fmt.Println("five")
	defer fmt.Println("six")

	// fmt.Println("Help! Something bad is happening.")
	// panic("Ending the program")
	// // fmt.Println("Waiting to execute")

	division(4, 2)
	division(8, 0)
	division(2, 8)

}

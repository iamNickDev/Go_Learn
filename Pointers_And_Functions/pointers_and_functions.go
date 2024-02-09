package main

import "fmt"

// function definition with a pointer argument
func update(num *int) {
	// dereference the pointer
	*num = 30
}

func display() *string {
	message := "Radha"

	return &message
}

func callByValue(num1 int) {
	num1 = 30
	fmt.Println(num1)
}
func callByReference(num2 *int) {
	*num2 = 10
	fmt.Println(*num2)
}
func main() {
	var number = 55

	var number1 int

	// function call
	update(&number)

	fmt.Println("The number is:", number)

	result := display()
	fmt.Println("Jai Shree", *result)

	callByValue(number1)
	callByReference(&number1)

}

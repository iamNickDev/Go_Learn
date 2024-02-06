package main

import "fmt"

var sum = 0

func findSquare(num int) int {
	square := num * num
	return square
}

// function that returns an anonymous function
func displayNumber() func() int {
	number := 10
	return func() int {
		number++
		return number
	}
}

func main() {
	// Anonymous Function
	var greet = func() {
		fmt.Println("Hello, how are you")
	}
	greet()

	// Anonymous Function with Parameters
	var sum = func(n1, n2 int) {
		sum := n1 + n2
		fmt.Println("Sum is: ", sum)
	}
	sum(5, 3)

	// Return Value From Anonymous Function
	var sumOf = func(n1, n2 int) int {
		sumOf := n1 + n2

		return sumOf
	}

	result := sumOf(10, 3)
	fmt.Println("SumOf is: ", result)

	// Return Value From Anonymous Function
	area := func(length, breadth int) int {
		return length * breadth
	}

	fmt.Println("Area of rectangle is: ", area(3, 4))

	// Anonymous Function as Arguments to Other Functions
	sum1 := func(number1 int, number2 int) int {
		return number1 + number2
	}
	result1 := findSquare(sum1(6, 9))
	fmt.Println("Result is: ", result1)

	a := displayNumber()
	fmt.Println(a())
}

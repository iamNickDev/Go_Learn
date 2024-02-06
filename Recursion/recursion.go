// Program to end the recursive function using if…else

package main

import (
	"fmt"
)

func countDown(number int) {

	if number > 0 {
		fmt.Println(number)

		// recursive call
		countDown(number - 1)
	} else {
		// ends the recursive function
		fmt.Println("Countdown Stops")
	}

}

func sum(number1 int) int {
	if number1 == 0 {
		return 0
	} else {
		return number1 + sum(number1-1)
	}
}

func factorial(numb int) int {
	if numb == 0 {
		return 1
	} else {
		return numb * factorial(numb-1)
	}
}

func main() {
	countDown(3)
	var num = 50
	var result = sum(num)
	fmt.Println("Sum:", result)

	var numb = 3

	var result1 = factorial(numb)
	fmt.Println("The factorial of 3 is:", result1)
}

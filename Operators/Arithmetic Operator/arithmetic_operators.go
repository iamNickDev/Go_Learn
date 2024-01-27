package main

import "fmt"

func main() {
	number1 := 7
	number2 := 2

	num1 := 15.0
	num2 := 4.0

	// + adds two variables
	sum := number1 + number2
	fmt.Printf("%d + %d = %d\n", number1, number2, sum)

	// - subtract two variables
	difference := number1 - number2
	fmt.Printf("%d - %d = %d\n", number1, number2, difference)

	// * multiply two variables
	product := number1 * number2
	fmt.Printf("%d * %d = %d\n", number1, number2, product)

	// / divide two integer variables
	quotient := number1 / number2
	fmt.Printf("%d / %d = %d\n", number1, number2, quotient)

	// / divide two floating point variables
	result := num1 / num2
	fmt.Printf("%g / %g = %g\n", num1, num2, result)

	// % modulo-divides two variables
	remainder := number1 % number2
	fmt.Println(remainder)

	// increment of num by 1
	number1++ //8
	fmt.Println(number1)

	// decrement of num by 1
	number1-- //7
	fmt.Println(number1)

}

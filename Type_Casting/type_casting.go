package main

import "fmt"

func main() {

	var number1 int = 20
	var number2 float32 = 5.7
	var sum float32

	var floatValue float32 = 5.45
	var integerValue int = 10

	// addition of different data types
	sum = float32(number1) + number2

	// type conversion from float to int
	var intValue int = int(floatValue)
	// type conversion from int to float
	var floatingValue float32 = float32(integerValue)

	fmt.Printf("Float Value is %g\n", floatValue)
	fmt.Printf("int Value is %d\n", intValue)

	fmt.Printf("integer Value is %d\n", integerValue)
	fmt.Printf("Float Value is %f\n", floatingValue)

	fmt.Printf("Sum is %g", sum)
}

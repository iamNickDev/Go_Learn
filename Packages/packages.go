// Golang main() package
package main

import (
	"Packages/Calculator/calculator.go"
	"fmt"
	"math"
	"strings"
)

// Commonly used packages in Go
// fmt Package
// math Package
// string Package

// Import package in Golang

func main() {

	var number int
	number = 10

	fmt.Println("Hello")

	// take input value
	fmt.Scan(&number)

	// print using Println
	fmt.Println("Number is", number)

	fmt.Print("Using Print")
	fmt.Println("Using Println")
	fmt.Scanf("%d", &number) // Input: 10
	fmt.Printf("%d", number) // Output: 10

	// find the square root
	fmt.Println(math.Sqrt(25)) // 5

	// find the cube root
	fmt.Println(math.Cbrt(27)) // 3

	// find the maximum number
	fmt.Println(math.Max(21, 18)) // 21

	// find the minimum number
	fmt.Println(math.Min(21, 18)) // 18

	// find the remainder
	fmt.Println(math.Mod(5, 2)) // 1

	// convert the string to lowercase
	lower := strings.ToLower("GOLANG STRINGS")
	fmt.Println(lower)

	// convert the string to uppercase
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

	// create a string array
	stringArray := []string{"I love", "Go Programming"}

	// join elements of array with space in between
	joinedString := strings.Join(stringArray, " ")
	fmt.Println(joinedString)

	number1 := 9
	number2 := 5

	// use the add function of calculator package
	fmt.Println(calculator.Add(number1, number2))

	// use the subtract function of calculator package
	fmt.Println(calculator.Subtract(number1, number2))

}

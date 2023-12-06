package main

import (
	"fmt"
	"math"
)

const s string = "constant "

// This is HelloWorld - main function
func main() {
	fmt.Println("Hello world")

	// Values
	//String
	fmt.Println("go" + "lang")

	//Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variables Declaration
	//Single
	var a = "initial"
	fmt.Println(a)

	// multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//Constants - Go supports constants of character, string, boolean, and numeric values.
	fmt.Print(s)

	const n = 5000000000

	const dec = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(dec))

	fmt.Println(math.Sin(n))
}

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

	// For Loop
	//The most basic type, with a single condition.

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//A classic initial/condition/after for loop.

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//for without a condition will loop repeatedly until you -
	// break out of the loop or return from the enclosing function.

	for {
		fmt.Println("loop")
		break
	}

	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// If/Else
	//Here’s a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//You can have an if statement without an else.
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	//Logical operators like && and || are often useful in conditions.

	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//A statement can precede conditionals; any variables declared in this
	// statement are available in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negetive")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

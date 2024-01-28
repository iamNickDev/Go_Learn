// We use the relational operators to compare two values or variables. For example,

// 5 == 6
// Here, == is a relational operator that checks if 5 is equal to 6.

// A relational operator returns

// true if the comparison between two values is correct
// false if the comparison is wrong
// Here's a list of various relational operators available in Go:

// Operator	Example	Descriptions
// == (equal to)	a == b	returns true if a and b are equal
// != (not equal to)	a != b	returns true if a and b are not equal
// > (greater than)	a > b	returns true if a is greater than b
// < (less than)	a < b	returns true if a is less than b
// >= (greater than or equal to)	a >= b	returns true if a is either greater than or equal to b
// <= (less than or equal to)	a <= b	returns true is a is either less than or equal to b

package main

import "fmt"

func main() {
	number1 := 12
	number2 := 20

	var result bool

	// equal to operator
	result = (number1 == number2)

	fmt.Printf("%d == %d return %t \n", number1, number2, result)

	// not equal to operator
	result = (number1 != number2)

	fmt.Printf("%d != %d returns %t \n", number1, number2, result)

	// greater than operator
	result = (number1 > number2)

	fmt.Printf("%d > %d returns %t \n", number1, number2, result)

	// less than operator
	result = (number1 < number2)

	fmt.Printf("%d < %d returns %t \n", number1, number2, result)
}

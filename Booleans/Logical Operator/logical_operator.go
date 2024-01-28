// Logical Operators in Go
// We use the logical operators to perform logical operations. A logical operator returns either true or false depending upon the conditions.

// Operator			Description		Example
// && (Logical AND)	exp1 && exp2	returns true if both expressions exp1 and exp2 are true
// || (Logical OR)	exp1 || exp2	returns true if any one of the expressions is true.
// ! (Logical NOT)	!exp			returns true if exp is false and returns false if exp is true.

// For More: https://www.programiz.com/golang/operators

package main

import "fmt"

func main() {
	number1 := 6
	number2 := 12
	number3 := 6

	var result bool

	// returns false because number1 > number2 is false
	result = (number1 > number2) && (number1 == number3)

	fmt.Printf("Result of AND operator is %t \n", result)

	// returns true because number1 == number3 is true
	result = (number1 > number2) || (number1 == number3)

	fmt.Printf("Result of OR operator is %t \n", result)

	// returns false because number1 == number3 is true
	result = !(number1 == number3)

	fmt.Printf("Result of NOT operator is %t \n", result)

}

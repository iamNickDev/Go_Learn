// Go while loop
package main

import "fmt"

func main() {
	number := 1
	multiplier := 1

	// Go while loop
	for number <= 5 {
		fmt.Println(number)
		number++
	}

	// multiplication table using while loop
	for multiplier <= 10 {
		product := 5 * multiplier

		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

	// Go do...while Loop
	for {
		// condition to terminate the loop
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++

	}
}

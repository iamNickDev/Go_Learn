package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		// terminates the loop when i is equal to 5
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// break-nested-loop Go break statement with nested loops
	for m := 1; m <= 3; m++ {
		for j := 1; j <= 3; j++ {
			if m == 2 {
				break
			}
			fmt.Println("\nm=", m, "\nj=", j)
		}
	}

	// Go continue statement
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("i=", i)
	}

	// Go continue statement with nested loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}

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

	for m := 1; m <= 3; m++ {
		for j := 1; j <= 3; j++ {
			if m == 2 {
				break
			}
			fmt.Println("\nm=", m, "\nj=", j)
		}
	}

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("i=", i)
	}
}

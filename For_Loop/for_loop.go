package main

import "fmt"

func main() {

	var n, sum = 10, 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}

//Go programming provides three different variations of the Scan() method:

// fmt.Scan()
// fmt.Scanln()
// fmt.Scanf()

package main

import (
	"fmt"
)

func main() {

	// create variables
	var name string
	var age int

	// take name and age input
	fmt.Println("Enter your name and age: \nNandkishore 24")
	fmt.Scan(&name, &age)

	// print input values
	fmt.Printf("Name: %s\nAge: %d", name, age)

}

//Go programming provides three different variations of the Scan() method:

// fmt.Scan()
// fmt.Scanln()
// fmt.Scanf()

package main

import "fmt"

func main() {

	// create variables
	var name string
	var age int
	var temperature float32
	var sunny bool

	// take name and age input
	fmt.Println("Enter your name and age: Nandkishore", 24)
	fmt.Scan(&name, &age)

	// print input values
	fmt.Printf("Name: %s\nAge: %d", name, age)

	fmt.Println("Enter your name and age: Nandkishore", 24)
	fmt.Scanln(&name, &age)
	fmt.Printf("Name: %s\nAge: %d", name, age)

	// take name and age input using format specifier
	fmt.Println("Enter your name and age: Nnadkishore ", 25)
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Name: %s\nAge: %d", name, age)
	// Take user input for temperature
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%f", &temperature)

	// Take user input for whether the day is sunny
	fmt.Println("Is the day sunny? (true/false):")
	fmt.Scanf("%t", &sunny)

	// Display the input values
	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)
}

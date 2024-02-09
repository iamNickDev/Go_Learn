package main

import "fmt"

type Weather struct {
	city        string
	temperature int
}

func main() {
	// Go Pointer to Struct
	// declare a struct Person
	type Person struct {
		name string
		age  int
	}

	person1 := Person{"nk", 24}

	// create a struct type pointer that
	// stores the address of person
	var ptr *Person
	ptr = &person1

	// print struct instance
	fmt.Println(person1)

	// print the struct type pointer
	fmt.Println(ptr)

	// access the name membe
	fmt.Println("Name:", ptr.name)
	// access the age member

	fmt.Println("Age:", ptr.age)
	// fmt.Println("Age:", (*ptr).age)

	// create struct variable
	weather := Weather{"Mathura", 20}
	fmt.Println("Initial Weather:", weather)

	// create struct type pointer
	ptr1 := &weather

	// change value of temperature to 25
	ptr1.temperature = 25

	fmt.Println("Updated Wether:", weather)

}

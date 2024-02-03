package main

import "fmt"

// initialize the function Rectangle
type Rectangle1 func(int, int) int

// create structure
type rectanglePara struct {
	length  int
	breadth int
	color   string

	// function as a field of struct
	rect1 Rectangle1
}

func main() {

	type Person struct {
		name string
		age  int
	}

	type Rectangle struct {
		length  int
		breadth int
	}

	person1 := Person{"Nk", 24}

	fmt.Println("Person1", person1)

	var person2 Person

	person2 = Person{
		name: "Radha",
		age:  1000000000,
	}

	fmt.Println("Person2", person2)

	fmt.Println("Person:", person1.name)

	rect := Rectangle{22, 12}

	fmt.Println("Length:", rect.length)
	fmt.Println("Breadth", rect.breadth)

	area := rect.length * rect.breadth

	fmt.Println("Area:", area)

	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect1: func(length int, breadth int) int {
			return length * breadth
		},
	}

	fmt.Println("Color of Rectangle:", result.color)
	fmt.Println("Area of Rectangle:", result.rect1(result.length, result.breadth))

}

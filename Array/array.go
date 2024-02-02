package main

import "fmt"

func main() {
	// declare array variable of type integer
	// defined size [size=5]
	var arrayOfInteger = [5]int{1, 2, 3, 4, 5}

	// Declare an array of undefined size
	// var arrayOfInteger := []int{1, 2, 3, 4, 5}
	// Shorthand notation to declare an array
	// arrayOfInteger := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arrayOfInteger)

	// Accessing array elements in Golang
	languages := []string{"golang", "Python", "Java", "Rust"}

	fmt.Println(languages)

	// access element at index 0
	fmt.Println(languages[0])

	// access element at index 2
	fmt.Println(languages[2])

	// Initialize an Array in Golang
	var arrayOfIntegers [3]int

	arrayOfIntegers[0] = 12
	arrayOfIntegers[1] = 30
	arrayOfIntegers[2] = 40

	fmt.Println(arrayOfIntegers)

	// Initialize specific elements of an array
	arrayOfInts := [5]int{0: 100, 3: 200}

	fmt.Println(arrayOfInts)

	// Changing the array element in Go
	weather := [3]string{"Rainy", "Sunny", "Cloudy"}

	weather[2] = "Stromy"
	fmt.Println(weather)

	// Find the length of an Array in Go
	length := len(weather)
	fmt.Println("The length of array is", length)

	// Looping through an array in Go
	age := [...]int{12, 4, 5}

	// loop through the array
	for i := 0; i < len(age); i++ {
	}

	// create a 2 dimensional array
	towDarray := [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(towDarray[i][j])
		}
	}
}

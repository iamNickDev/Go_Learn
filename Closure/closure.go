package main

import "fmt"

func greet(name string) {
	var displayName = func() {
		fmt.Println("Hi", name)
	}
	displayName()
}

func greet1() func() {
	return func() {
		fmt.Println("Shree Radhaji")
	}
}

func testCloser() func() string {
	name := "Shree Radhaji"

	return func() string {
		name = "Jai " + name
		return name
	}
}

func calculate() func() int {
	num := 1

	return func() int {
		num = num + 2
		return num
	}
}

func displayNumber() func() int {
	number := 0

	return func() int {
		number++
		return number
	}
}

func main() {
	greet("radhaji")
	g1 := greet1()
	g1()

	message := testCloser()

	fmt.Println(message())

	odd := calculate()

	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

	num1 := displayNumber()

	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())

	num2 := displayNumber()

	fmt.Println(num2())
	fmt.Println(num2())

}

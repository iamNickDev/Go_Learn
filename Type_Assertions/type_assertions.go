package main

import "fmt"

func main() {
	var a interface{}

	a = 10

	interfaceValue, booleanValue := a.(int)

	fmt.Println("Interface Value:", interfaceValue)
	fmt.Println("Boolean Value:", booleanValue)
}

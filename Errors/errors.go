package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type DivisionByZero struct {
	message string
}

func (z DivisionByZero) Error() string {
	return "Number Cannot Be Divided by Zero"
}

func divide1(n1 int, n2 int) (int, error) {
	if n2 == 0 {
		return 0, &DivisionByZero{}
	} else {
		return n1 / n2, nil
	}
}

func checkName(name string) error {
	newError := errors.New("Invalid Name")

	if name != "Kanha" {
		return newError
	}
	return nil
}

func divide(num1, num2 int) error {
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero\n", num1, num2)
	}
	return nil
}

func main() {
	name := "Hello"

	err := checkName(name)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Name")
	}

	err2 := divide(4, 0)

	if err != nil {
		fmt.Printf("Error: %s", err2)
	} else {
		fmt.Println("Valid Division")
	}

	number1 := 15

	number2 := 0

	result, err := divide1(number1, number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
}

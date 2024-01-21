package main

import "fmt"

type Shape interface {
	area() float32
}

type Rectangle struct {
	length, breadth float32
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func calculate(s Shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	rect := Rectangle{7, 4}

	calculate(rect)
}

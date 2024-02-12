package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface
func (t Tringle) area() float32 {
	return 0.5 * t.base * t.height
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

type Tringle struct {
	base, height float32
}

// access method of the interface
func calculate(s Shape) float32 {
	return s.area()
}

func displayValue(i ...interface{}) {
	fmt.Println(i...)
}

// main function
func main() {

	var a interface{}
	fmt.Println("Value: ", a)

	j := "welcome to Bharat"
	b := 20
	c := false

	displayValue(j)
	displayValue(j, b)
	displayValue(j, b, c)

	r := Rectangle{7, 4}
	t := Tringle{8, 12}

	rectangleArea := calculate(r)
	fmt.Println("Area of Rectangle:", rectangleArea)

	triangleArea := calculate(t)
	fmt.Println("Area of Triangle:", triangleArea)
}

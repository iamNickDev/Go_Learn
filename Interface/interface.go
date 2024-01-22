package main

import "fmt"

type Shape interface {
	area() float32
}

type Rectangle struct {
	length, breadth float32
}

func (t Tringle) area() float32 {
	return 0.5 * t.base * t.height
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

type Tringle struct {
	base, height float32
}

func calculate(s Shape) float32 {
	return s.area()
}

func main() {

	r := Rectangle{7, 4}
	t := Tringle{8, 12}

	rectangleArea := calculate(r)
	fmt.Println("Area of Rectangle:", rectangleArea)

	triangleArea := calculate(t)
	fmt.Println("Area of Triangle:", triangleArea)
}

package main

import (
	"fmt"
)

// Constants declared outside of function
const NAME string = "notChangeble" //Typedconstants
const B = 2                        // Untyped constants

// Multiple Constants Declaration
const (
	SAM     int = 1
	BAM         = 2
	CAM         = "ram"
	POINTER     = 3.12
)

func main() {
	fmt.Println("Hello World!")

	//	Go Variables
	var studentName string = "nick"
	var i, j string = "All", "well"
	studentAge := 24 //"short variable declaration" or "short assignment statement."
	var a, b, c, d int = 1, 2, 3, 4
	var e, f, g, h = 8, "ram", "shyam", false

	var m string = "myName"
	var n int = 15
	var p int = 10
	pointerValue := 12.5

	var (
		MyNum    int             //Pascal Case
		myNumber int    = 1      //Camel Case
		my_name  string = "nick" //Snake Case
	)

	//Constants declared inside of function

	const A int = 1 // Typedconstants
	const PI = 3.12 // Untyped constants

	fmt.Println(studentName)
	fmt.Println(i, j)
	fmt.Printf("m has a value: %v and type: %T\n", m, m)
	fmt.Printf("p has a value: %#p and type: %T\n", p, p)
	fmt.Printf("n has a value: %v and type: %T\n", n, n)
	fmt.Printf("pointerValue has a value: %v and type: %T", pointerValue, pointerValue)

	fmt.Println(studentAge)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(myNumber)
	fmt.Println(my_name)
	fmt.Println(MyNum)
	fmt.Println(NAME)
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(SAM)
	fmt.Println(BAM)
	fmt.Println(CAM)
	fmt.Println(POINTER)
}

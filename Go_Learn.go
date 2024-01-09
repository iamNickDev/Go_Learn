package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

const s string = "constant "

// This is HelloWorld - main function
func main() {
	fmt.Println("Hello world")

	//----------------------------------------------------------------*********------------------------------------------------------------------

	// Values
	//String
	fmt.Println("go" + "lang")

	//Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variables Declaration
	//Single
	var a = "initial"
	fmt.Println(a)

	// multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//----------------------------------------------------------------*********------------------------------------------------------------------

	//Constants - Go supports constants of character, string, boolean, and numeric values.
	fmt.Print(s)

	const n = 5000000000

	const dec = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(dec))

	fmt.Println(math.Sin(n))

	//----------------------------------------------------------------*********------------------------------------------------------------------

	// For Loop
	//The most basic type, with a single condition.

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//A classic initial/condition/after for loop.

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//for without a condition will loop repeatedly until you -
	// break out of the loop or return from the enclosing function.

	for {
		fmt.Println("loop")
		break
	}

	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//----------------------------------------------------------------*********------------------------------------------------------------------

	// If/Else
	//Here’s a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//You can have an if statement without an else.
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	//Logical operators like && and || are often useful in conditions.

	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//A statement can precede conditionals; any variables declared in this
	// statement are available in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negetive")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//----------------------------------------------------------------*********------------------------------------------------------------------

	// Switch Statements - Switch statements express conditionals across many branches.

	// basic switch.
	g := 2
	fmt.Print("write ", g, " as ")
	switch g {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//You can use commas to separate multiple expressions in the same case statement.
	// We use the optional default case in this example as well.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weekend")
	default:
		fmt.Println("it's a weekday")
	}

	//switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("It's a afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	//----------------------------------------------------------------*********------------------------------------------------------------------

	//Arrays - In Go, an array is a numbered sequence of elements of a specific length.
	//In typical Go code, slices are much more common; arrays are useful in some special scenarios.

	var arr [5]int
	fmt.Println("emp:", arr)

	//We can set a value at an index using the
	//array[index] = value syntax, and get a value with array[index].
	arr[4] = 100
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	//The builtin len returns the length of an array.
	fmt.Println("len:", arr)

	//Use this syntax to declare and initialize an array in one line.
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	//Array types are one-dimensional, but you can compose types
	//to build multi-dimensional data structures.

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//----------------------------------------------------------------*********------------------------------------------------------------------

	//Slices - Slices are an important data type in Go,
	//giving a more powerful interface to sequences than arrays.

	//Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//To create an empty slice with non-zero length,
	//use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
	//By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time,
	//it’s possible to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	//We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("grt:", s[2])

	//len returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	//slices support several more that make them richer than arrays.
	//One is the builtin append, which returns a slice containing one or more new values.
	//Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("spt:", s)

	//Slices can also be copy’d.
	//Here we create an empty slice c of the same length as s and copy into c from s.
	copyLen := make([]string, len(s))
	copy(copyLen, s)
	fmt.Println("cpy:", c)

	//Slices support a “slice” operator with the syntax slice[low:high].
	//For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5].
	l2 := s[:5]
	fmt.Println("sl2:", l2)

	// And this slices up from (and including) s[2].
	l3 := s[2:]
	fmt.Println("sl3:", l3)

	//We can declare and initialize a variable for slice in a single line as well.
	singleLine := []string{"g", "h", "i", "j"}
	fmt.Println("singleLine:", singleLine)

	singleLine2 := []string{"g", "h", "i", "j"}
	if slices.Equal(singleLine, singleLine2) {
		fmt.Println("singleLine == singleLine2")
	}

	//Slices can be composed into multi-dimensional data structures.
	//The length of the inner slices can vary,
	//unlike with multi-dimensional arrays.
	towD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", towD)

}

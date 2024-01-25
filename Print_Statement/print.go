// We use these three functions to print output messages in Go programming.

// fmt.Print()
// fmt.Println()
// fmt.Printf()

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Radha\n")
	fmt.Print("Radha!\n")

	name := "Radha\n"
	fmt.Print(name)
	fmt.Print("Name:", name)

	naam := "Radha"
	fmt.Println("Suddh Naam:", naam)

	var ladliju = "Radhaji"
	lilaye := 1000000000

	fmt.Printf("%s did %d of lilas with lalju ", ladliju, lilaye)

	println("\njai shree krishna")
	println(10)

}

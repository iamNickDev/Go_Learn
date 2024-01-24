// We use these three functions to print output messages in Go programming.

// fmt.Print()
// fmt.Println()
// fmt.Printf()

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Radha, ")
	fmt.Print("Radha!\n")

	name := "Radha"
	fmt.Print(name)
	fmt.Print("\nName:", name)

	naam := "\nRadha"
	fmt.Println(naam)

}

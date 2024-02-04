package main

import (
	"fmt"
	"strings"
)

func main() {
	// Golang String
	var message1 = "Hari Om"
	message2 := "Shree Radha Radha"

	// Golang String using backtick (` `)
	message := `Jai Shree Radha`

	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"
	text1 := "go is fun to learn"
	text2 := "I LOVE GOLANG"

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message)

	// Access Characters of String in Go
	fmt.Printf("%c\n", message[0])
	fmt.Printf("%c\n", message[4])
	fmt.Printf("%c\n", message[10])

	// Find the length of a string
	fmt.Println("Length of a string is: ", len(message))

	// Join Two Strings Together
	result := message1 + " " + message2
	fmt.Println("Merged together: ", result)

	// Compare Two Strings in Go
	// to compare two strings: string1 and string2. The function returns:
	// -1 because string1 is smaller than string2
	// 1 because string2 is greater than string3
	// 0 because string1 and string3 are equal
	fmt.Println(strings.Compare(message1, message2))

	// Check if String contains Substring
	conclusion := strings.Contains(text, substring1)
	conclusion2 := strings.Contains(text, substring2)
	fmt.Println("Result:", conclusion)
	fmt.Println("Result Substring:", conclusion2)

	// Replace a String in Go
	replaceString := strings.Replace(text, "Go", "go", 1)
	fmt.Println(replaceString)

	// Change Case of Go String
	// convert to uppercase
	text1 = strings.ToUpper(text1)
	fmt.Println(text1)

	// convert to lowercase
	text2 = strings.ToLower(text2)
	fmt.Println(text2)

	// Split Strings in Golang
	splittedString := strings.Split(text1, " ")
	fmt.Println(splittedString)

}

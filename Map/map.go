package main

import (
	"fmt"
)

func main() {
	subjectMarks := map[string]float32{"GoLang": 85, "Java": 78, "Python": 80}
	fmt.Println(subjectMarks)

	fmt.Println(subjectMarks["GoLang"])
	fmt.Println(subjectMarks["Java"])
	fmt.Println(subjectMarks["Python"])

	subjectMarks["Java"] = 40

	fmt.Println("Updated Map:", subjectMarks)

	subjectMarks["Solidity"] = 60

	fmt.Println("New element added:", subjectMarks)

	delete(subjectMarks, "Python")

	fmt.Println("Element deleted:", subjectMarks)

	for subject, marks := range subjectMarks {
		fmt.Printf("Marks of %s is %v\n", subject, marks)
	}
}

package main

import "fmt"

func main() {
	dayOfWeek := 4

	dayOfWeek2 := "Sunday"

	numberOfDays := 28

	// Program to print the day of the week using fallthrough in switch

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")
		fallthrough

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")

	}

	// Program to check if the day is a weekend or a weekday

	switch dayOfWeek2 {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}

	// Program to check if it's February or not using switch without expression
	switch {
	case 28 == numberOfDays:
		fmt.Println("it's February")

	default:
		fmt.Println("Not February")
	}

	// Program to check the day of a week using optional statement

	switch day := 4; day {

	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid Day!")
	}
}

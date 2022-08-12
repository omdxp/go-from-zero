package main

import (
	"fmt"
	"time"
)

func main() {
	language := "Go"

	switch language {
	case "Python":
		fmt.Println("You selected Python")
		// break // compile time warning "redundant break statement"
	case "Go", "Golang":
		fmt.Println("Wohoo! you selected Go!")
	default:
		fmt.Println("You have selected other languages")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even")
	case n%2 != 0:
		fmt.Println("Odd")
	}

	hour := time.Now().Hour()
	switch { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// in a switch statement, you can use curly braces in each case if you want more readability
// by default indentation will work
// also the compared cases must be of the same type as the compared value
// Go adds break statement at the end of each case statement automatically
// default clause is not mandatory
// expressions are evaluated as well

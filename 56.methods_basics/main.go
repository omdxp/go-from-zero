package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() { // n is the actual copy of names that are being used (like self or this for OOP)
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds() // this is a method
	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"Omar", "Bea"}
	friends.print()

	// another way to do it
	names.print(friends) // using print as a function

	var n int64 = 533253
	fmt.Println(n)
	fmt.Println(time.Duration(n))
}

// Go does not have classes, but methods can be defined on predefined types
// a type may have method set associated with it which enhances the type with extra behavior
// methods are also known as receiver functions
// it is good practice to use something short and concise of the receiver parameter
// a method belongs to a type, a function belongs to a package
// it is idiomatic in Go to convert a type to an expression to access a specific method

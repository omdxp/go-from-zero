package main

import "fmt"

// function that returns a function that returns an int
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("Hello")

	a := increment(10) // first class function
	fmt.Printf("%T\n", a)
	a()              // +1
	fmt.Println(a()) // +1
}

// anonymous function is a function that don't have a name and is declared inline using function literal
// anonymous functions can be used as closures
// Go supports first class functions, means functions can be used as variables and can be passed as arguments and returned from from other functions

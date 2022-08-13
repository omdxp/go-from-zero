package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function!")
}

func f2(a int, b int) { // a and b are local to f2
	fmt.Println("Sum:", a+b)
}

// shorthand parameters notation
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// return type
func f4(a float64) float64 {
	return math.Pow(a, a)
	// any statement after are never executed
}

// multiple return statement
func f5(a, b int) (int, int) {
	return a + b, a * b
}

// named return values
func sum(a, b int) (s int) { // behind the scene, Go declare var s int
	fmt.Println(s)
	s = a + b
	return // naked return statement (s will be returned since it's the expected return value)
}

func main() {
	f1() // calling f1
	f2(5, 7)
	f3(4, 5, 6, 4.4, 7.8, "Golang")
	p := f4(5.1)
	fmt.Println(p)
	// multiple assignment
	a, b := f5(8, 9) // if one of the values is not needed "_" is used so the compiler won't complain about it
	fmt.Println(a, b)

	mySum := sum(4, 8)
	fmt.Println(mySum)
}

// everything in Go is passed by value (there's no passing by reference)
// if a function return multiple values and error, by convention the last one should the error

package main

import "fmt"

func main() {
	const a float64 = 5.1 // typed constant

	const b = 6.7 // untyped constant

	const c float64 = a * b
	const str = "Hello " + "Go!"

	const d = 5 > 10
	fmt.Println(d)

	// for typed constants you cannot perform operations on different types
	// const x int = 5
	// const y float64 = 2.2 * x

	// for untyped constans you can
	const x = 5
	const y = 2.2 * 5
	fmt.Printf("%T\n", y)

	var i int = x     // x will change to int
	var j float64 = x // x will change to float64 (var j float64 = float64(x))
	var p byte = x    // x will change to byte
	fmt.Println(i, j, p)

	const r = 5 // default type is untyped int
	var rr = r  // rr type is int
	fmt.Printf("%T\n", rr)
}

package main

import "fmt"

func main() {
	var numbers [4]int // elements are initialized with zeros

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Omar", "Diana", "Bea", "Mathew"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"} // only 2 first elements are initialzed, the rest with zero values ("")
	fmt.Printf("%#v\n", a4)

	// ellipsis operator ...
	a5 := [...]int{1, 2, 3, 4, 5} // length can be deducted during initialization
	fmt.Printf("%#v\n", a5)
	fmt.Printf("length of a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3, // this comma is mandatory
	}
	fmt.Println(a6)
}

// arrays in Go are fixed length

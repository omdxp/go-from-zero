package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	// fmt.Println(a == b) // compile time error, slices can only be compared to nil

	var eq bool = true
	for i, va := range a {
		if va != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) { // this is very important when comparing slices
		eq = false
	}

	fmt.Println("a == b?", eq)
}

// slices can only be compared by nil or by iterating on all elements

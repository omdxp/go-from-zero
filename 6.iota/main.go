package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota
		East
		South
		West
	)
	fmt.Println(West)

	const (
		a = (iota * 2) + 1 // 1
		b                  // 3
		c                  // 5
	)
	fmt.Println(a, b, c)

	const (
		x = -(iota + 2) // -2
		_               // -3 (blank identifiers are skipped)
		y               // - 4
		z               // -5
	)
	fmt.Println(x, y, z)
}

// iota is number generator for constants
// by it starts by 0 and increment automatically
// the initial value, also the increment step can be changed

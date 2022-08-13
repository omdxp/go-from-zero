package main

import "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

func changeVar(a int) {
	a = 66
}

func main() {
	x := 8
	p := &x

	fmt.Println("value of x before calling change():", x)
	change(p) // this will change the value of x
	fmt.Println("value of x after calling change():", x)

	fmt.Println("value of x before calling changeVar():", x)
	changeVar(x) // this will not change the value of x
	fmt.Println("value of x after calling changeVar():", x)
}

// functions create a copy of pointers for the arguments that share the same address
// in Go, functions always works on copies and not originals
// everything is passed by value

package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)

	fmt.Printf("*p1 is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)
	fmt.Printf("**pp1 is %v\n", **pp1)

	**pp1++ // increments a
	fmt.Printf("a is %v\n", a)

	// comparing pointers
	var p2 *int
	fmt.Printf("%#v\n", p2)

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) // false

	p4 := &y
	fmt.Println(p2 == p4) // true
}

// 2 pointers are equal if and only if they point to the same variable or they are both nil

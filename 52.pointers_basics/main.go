package main

import "fmt"

func main() {
	name := "Omar"
	fmt.Println(&name) // & is the address of operator

	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	// declare a pointer without intilializing it
	var ptr1 *float64
	_ = ptr1

	// another way to create a pointer
	p := new(int) // p is a pointer to int (*int)

	x = 100
	p = &x

	fmt.Printf("p is of type %T with a value of %v\n", p, p)
	fmt.Printf("address of x is %p\n", &x)

	// derefrence operator
	*p = 90 // will change also the x value because it changes the value in that address (same as x = 90)
	fmt.Println(x, *p)
	fmt.Println("*p==x:", *p == x)
	fmt.Println("p==&x:", p == &x)

	*p = 10     // x = 10
	*p = *p / 2 // x / 2
	fmt.Println(x)

	// &value => pointer
	// *pointer => value
}

// variables are nothing but alphanumeric nickname for values stored in some memory addresses
// pointers hold the address memory of variables
// pointers are initialized to nil (their zero value is nil)
// unlike C, Go does not support pointer arithmitics (like incrementing pointers) it'll throw a compile error if doing so
// a pointer itself is also a variable stored somewhere in memory
// new() builtin function takes the type as argument, allocate enough memory to accommodate the value of that type, and returns a pointer to it

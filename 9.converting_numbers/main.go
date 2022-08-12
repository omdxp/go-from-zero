package main

import "fmt"

func main() {
	var x = 3   // int type
	var y = 3.1 // float type

	// x = x * y // compile time error because 3 is unnamed constant untyped int and 3.1 is unnamed constant untyped float
	x = x * int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)

	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5 // int type
	fmt.Printf("%T\n", a)

	var b int64 = 2
	a = int(b)
}

// int(y) will return a new int value, and does not mutate the original variable

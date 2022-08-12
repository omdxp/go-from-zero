package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hour
	fmt.Printf("duration in seconds: %v\n", duration*secondsInHour)

	x, y := 5, 0
	fmt.Println(x / y) // this error is discovered at runtime (panic error)

	const a, b = 5, 0
	fmt.Println(a / b) // this error is discovered at compile time

	// multiple declaration of constants
	const (
		min1 = -500
		min2 // -500
		min3 // -500
	)
}

// all variable literals are "unnamed constants" (like 3.14, true, "hello")
// constants values are known at compile time
// the compiler warns about unused constants, where it shows error for unused variables
// constants must be initialized when declared
// constants rules:
//	- cannot modify constant
// 	- cannot initialize a constant at runtime (like using function calls that belongs to runtime "const p = math.Pow(2, 3)")
//	- cannot use a variable to initialize a constant
//	- can use a len() function on constants to initialize a constant (like using it on string literals because they're unnamed constants "cont l = len("hello")") (bultin functions are available at compile time)

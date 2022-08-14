package main

import "fmt"

// declare an array
var numbers [10]int

func init() {
	fmt.Println("This is init #1")
}

func init() {
	fmt.Println("This is init #2")
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

func main() {
	fmt.Println("This is main")
	// init() // compile time error, undecclared name: init
	fmt.Printf("%#v\n", numbers)
}

// in Go there 2 functions reserved for special purposes: main and init
// main and init do not take arguments or return anything
// main function is the starting point to execute the program
// init function is called automatically when the package is initialized and it is always called before main function
// int function can be defined multiple times in a file or a package and the order of execution will be according to the order where they appear
// purpose of init is to initialize global variables that cannot be initialized otherwise in the global context

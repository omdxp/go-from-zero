package main

import "fmt"

type emptyInterface interface{}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{4, 5, 6}
	fmt.Println(empty)
	// fmt.Println(len(empty)) // compile time error, invalid argument for len()
	fmt.Println(len(empty.([]int))) // type assertions must be used

	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5.6, 5., 7.8}
	fmt.Println(you.info)
}

// any type satisfies an empty interface and that means it can represents any value
// fmt.Println() function uses a slice of empty interfaces
// in fact in a lot of Go specifications, empty interfaces are used
// empty interfaces are not used directly in operations
// interfaces stores 2 values, a dynamic type and a dynamic concrete value
// empty interfaces are used in a code that handles values of unknown types
// empty interfaces can cause the program to be hard to maintain
// empty interfaces are used only when it is necessary

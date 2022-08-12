package main

import "fmt"

func main() {
	var cities []string // nil
	fmt.Println(cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities)) // 0

	// cities[0] = "Berlin" // runtime error, index out of range

	numbers := []int{2, 3, 4, 5} // slice literal
	fmt.Println(numbers)

	nums := make([]int, 2) // initialize a slice with size of 2 with zero values
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Omar", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println(myFriend)

	friends[1] = "Bea"
	fmt.Println(friends)

	// iterate over a slice
	for i, v := range numbers {
		fmt.Printf("index %d value %v\n", i, v)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}

// unlike arrays, slices can shrink and grows at runtime, they don't have fixed length
// by default slice is initialized with nil

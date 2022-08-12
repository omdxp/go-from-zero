package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	// numbers[5] = 100 // compile time error, index out of bounds

	// iterate through the array
	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}

	// multidimensional array
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10}, // [3]int is not mandatory
	}
	fmt.Println(balances)

	// the following array are different, they're not connected in anyway
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	m[0] = -1
	fmt.Println("n is equal to m:", n == m)
}

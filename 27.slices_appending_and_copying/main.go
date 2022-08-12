package main

import "fmt"

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 10)
	fmt.Println(numbers)

	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...) // ellipsis operator is used to spread all elements
	fmt.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, 2) // len(src) to make sure to copy all elements
	nn := copy(dst, src)  // in this case it'll copy only the first 2 elements
	fmt.Println(src, dst, nn)
}

// append() bultin function returns a new slice and do not mutate the original
// copy() builtin function copies elements from source slice to destination slice and returns number of elements that have been copied
// copy() copies only the elements for the available destination slice

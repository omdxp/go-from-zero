package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) { // int slice
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(a ...float64) (float64, float64) {
	sum, product := .0, 1.
	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

// mix non-variadic with variadic parameters
func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	return fmt.Sprintf("Age: %d, Full Name: %s\n", age, fullName)
}

func main() {
	f1(1, 2, 3, 4)
	f1()

	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)
	_ = nums

	f1(nums...) // pass slice elements to variadic function

	f2(nums...)
	fmt.Println("Nums:", nums) // nums have been modified in f2() function

	s, p := sumAndProduct(2, 5., 10., 5.6, 5.6, 87.3)
	fmt.Println(s, p)

	info := personInfo(24, "Omar", "Blg")
	fmt.Println(info)
}

// a variadic function is a function that can accept as many parameter as you want
// fmt.Println() is a variadic function
// append() builtin function is a variadic function too

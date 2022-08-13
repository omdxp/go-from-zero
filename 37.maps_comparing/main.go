package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}

	// fmt.Println(a == b) // compile time error, maps can only be compared to nil

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}

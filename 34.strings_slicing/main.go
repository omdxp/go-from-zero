package main

import "fmt"

func main() {
	s1 := "I love Golang!"
	fmt.Println(s1[2:6])

	s2 := "ꑕàéç"
	fmt.Println(s2[1:3]) // it returns bytes (strange characters ��)

	// to slice better a string we convert it to slice of runes first
	rs := []rune(s2)
	fmt.Printf("%T\n", rs)

	// the second step is to slice the runes slice and convert it back to string
	fmt.Println(string(rs[1:3]))
}

// converting between different types (like between slice of runes and strings) is not efficient, a new backing array is created each time

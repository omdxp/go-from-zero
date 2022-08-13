package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1)) // since s1 contains only ASCII characters it will return 6

	name := "àstro"
	fmt.Println(len(name)) // since name has à (not ASCII character encoded by 2 bytes) it will return 6

	fmt.Println(len("ꑕ")) // this rune occupies 3 bytes

	// calculate runes instead of bytes
	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("ꑕ")
	fmt.Println(n, m)
}

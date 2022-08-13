package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b'

	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "àstro"
	fmt.Println(len(str)) // 6 (because the number of bytes à is 2)

	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println("\n" + strings.Repeat("#", 20))
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println("\n" + strings.Repeat("#", 20))
	for _, r := range str {
		fmt.Printf("%c", r)
	}

}

// rune literals are declared with '' (signle quotes) and they're represented as int32 code points
// rune is an alias for int32
// len() builtin function return the number of bytes in a string
// when using indexing a string we get the byte positions not runes
// when iteration with foor loop with incrementing i by 1, it will go through string byte by byte
// range will iterate a string decoded with runes

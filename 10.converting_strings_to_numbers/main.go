package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99) // this converts the ASCII value of 99 which is 'c'
	fmt.Println(s)

	// s1 := string(44.2) // compile time error
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 42069)
	fmt.Println(myStr1)

	fmt.Println(string(42069))

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1 * 3.4)

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2)
}

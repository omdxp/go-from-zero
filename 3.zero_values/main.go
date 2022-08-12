package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	a = int(b)
	fmt.Println(a, b)

	var (
		value   int
		price   float64
		name    string
		done    bool
		pointer *any
	)
	fmt.Println(value, price, name, done, pointer)
}

// Go is statically type language, variable types are known at compile time
// however, the values are known at runtime
// in Go all variables are initialized
//	- numeric types: 0
// 	- bool types: false
// 	- string types: "" (empty string)
// 	- pointer types: nil

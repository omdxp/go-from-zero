package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte = a // byte is an alias for uint8
	_ = b

	type second = uint // second is an alias for uint

	var hour second = 3600
	fmt.Printf("minutes in 1 hour: %d\n", hour/60)
}

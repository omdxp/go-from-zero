package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ // overflow, x is 0
	fmt.Println(x)

	// a := int8(255 + 1) // expressions are evaluated at compile time (here overflow is catched at compile time)

	var b int8 = 127
	fmt.Printf("%d\n", b+1) // it sets it to the minimum value

	b = -128
	b-- // it sets it to the maximum value
	fmt.Println(b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f = f * 1.2
	fmt.Println(f)

	// const i int8 = 300 // compile time error
}

// integers overflows at runtime between their min and max values
// floats overflows to +Inf or -Inf
// "math/big" is useful for big numbers

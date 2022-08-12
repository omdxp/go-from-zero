package main

import "fmt"

func main() {
	// int type
	var i1 int8 = -128 // 8 bits to represent this integer (min -128, max 127) if it overwlows we'll get a compile time error
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65_535 // unsigned int represented in 16 bits
	fmt.Printf("%T\n", i2)

	// float type
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// complex type
	var c1, c2 complex64 = complex(1, 1), complex(0, -2)
	fmt.Printf("%T %T\n", c1, c2)

	// rune type (used for characters)
	var r rune = 'f'
	fmt.Printf("%T %v\n", r, r)

	// bool type
	var b bool = true
	fmt.Printf("%T\n", b)

	// string type (sequence of bytes)
	var s string = "Hello Go ✋🏻!"
	fmt.Printf("%T\n", s)

	// composite types
	// array type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// slice type
	var cities = []string{"San Francisco", "Berlin", "Tokyo"}
	fmt.Printf("%T\n", cities)

	// map type
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// struct type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	// pointer type
	var x int = 2
	ptr := &x
	fmt.Printf("%T\n", ptr)

	// function type
	fmt.Printf("%T\n", f)

	// interface type
	type Human interface{}

	// channel type
	c := make(chan int)
	fmt.Printf("%T\n", c)
}

func f() {}

// int is an alias for int32 or int64 based on the platform
// uint is an alias for uint32 or uint64 based on the platform
// byte is an alias for uint8
// rune is an alias for int32

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("Number of arguments in os.Args:", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T %T\n", os.Args[1], result)
	_ = err
}

// arguments are always strings (os.Args is a slice of strings)

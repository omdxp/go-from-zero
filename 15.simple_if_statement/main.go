package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45s")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("20x"); err == nil {
		fmt.Println("No error, i is", i)
	} else {
		// handle error
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be integer! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
	}
}

// if initialization; condition {}

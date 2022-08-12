package main

import "fmt"

func main() {
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	goto todo // compile time error, cannot jump over variable decalaration
	x := 5
todo:
	fmt.Println("Something here")
}

// the use of goto in any language is not recommended as it will be difficult to trace the control flow of the program
// goto statements are dangerous and must be used with caution

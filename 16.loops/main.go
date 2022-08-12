package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// while loop
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// infinite loop
	sum := 0
	for {
		sum++
	}
	fmt.Println(sum) // this line is never reached (detected at compile time)
}

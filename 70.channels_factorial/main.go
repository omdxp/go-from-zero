package main

import "fmt"

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	c <- f
}

func main() {
	ch := make(chan int)
	defer close(ch) // best practice

	go factorial(5, ch)

	// blocking call
	f := <-ch
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	for i := 5; i <= 15; i++ {
		go func(i int) {
			factorial(i, ch)
		}(i)
		f := <-ch
		fmt.Println(f)
	}
}

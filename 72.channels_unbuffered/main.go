package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) // unbuffered channel

	c2 := make(chan int, 3) // buffered channel (with capacity of 3)
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10                                                           // this will block main goroutine until it wakes up (when main receive the value)
		fmt.Println("func goroutine after sending data into the channel") // this will exectute after main receives the value
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(2 * time.Second)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	time.Sleep(time.Second)
}

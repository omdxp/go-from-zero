package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / 1_000_000
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Csec!"
	}()

	time.Sleep(2 * time.Second) // force waiting the channels to be not empty

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("Received:", msg)
		case msg := <-c2:
			fmt.Println("Received:", msg)
		default: // this when none of the channels have been received
			fmt.Println("no activity")
		}
	}

	end := time.Now().UnixNano() / 1_000_000

	fmt.Println(end - start)
}

// select statement let a goroutine wait on multiple communication operations
// it is similar to switch statement but is for channels
// it execute a case when that channel is ready to be received from
// it chooses one case at random if multiple cases are ready
// select is used to wait on multiple goroutines simultaneously
// default case in select is called "none blocking operation"

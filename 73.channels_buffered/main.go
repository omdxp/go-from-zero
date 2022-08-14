package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // buffered channel with capacity of 3

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c)

	time.Sleep(2 * time.Second)

	for v := range c { // v := <- c
		fmt.Println("main goroutine received value from channel:", v)
	}

	fmt.Println(<-c) // 0 (zero value of int type) because c is closed
	// c <- 10 // panic: send on closed channel
}

// the sender of a buffered channel will block only when there is no empty slot in the channel while the receiver will block on the channel when its empty
// you don't need to close every channel when you are done with it
// it is only necessary to close a channel when it is important to tell the receiver goroutines that all data have been sent
// communication over an unbuffered channel causes the sending and receiving goroutines to synchronize
// unbuffered channels are called "synchronous channels"
// if you try to send into a closed channel it'll panic!
// in buffered channels the sending and receiving operations are decoupled and not synchronized
// channel buffering may also affect program performance

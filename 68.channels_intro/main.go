package main

import "fmt"

func main() {
	var ch chan int // unintialized channel takes zero value which is nil
	fmt.Println(ch)

	ch = make(chan int) // initialize the channel
	fmt.Println(ch)     // address of channel

	// short declaration operator (declare and initialize)
	c := make(chan int)

	// <- channel operator

	// send operation
	c <- 10

	// receive operation
	num := <-c
	fmt.Println(<-c)
	_ = num

	// close a channel
	close(c)
}

// channel are used to communicate between goroutines
// a channel stores an address, it's like a pointer
// passing channels to functions has the same effect as passing pointers to functions
// close() builtin function closes a channel
// channels is used with conjunction with goroutines, if there are no goroutines started there will be a fata error: deadlock!

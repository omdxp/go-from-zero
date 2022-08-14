package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n // sending n into ch
}

func main() {
	// bidirectional channel
	c := make(chan int) // this channel will send and receive values of type int

	// only for receiving
	c1 := make(<-chan string) // this channel will only receive values of type string

	// only for sending
	c2 := make(chan<- string) // this channel will only send values of type string

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)

	n := <-c // recieve value into n (main will wait until there is a value in c)
	// if there is no value received from c, there will be a deadlock!
	fmt.Println("Value received:", n)
	fmt.Println("Exiting main...")
}

// there bidirectional channels (send and receive) and uniderctional channels (can only send or receive)

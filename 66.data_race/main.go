package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100 // 100 goroutines

	var wg sync.WaitGroup
	wg.Add(gr * 2) // 100 goroutines will increment, and 100 goroutines will decrement a shared value

	var n int = 0 // the shared value

	for i := 0; i < gr; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			n++
		}()

		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			n--
		}()
	}

	wg.Wait()
	fmt.Println(n) // different values of n because of the data race
}

// data race is among the major problems in concurrent systems
// data race occurs because of unsynchronized access to shared memory
// data race can be check when running the program by the -race flag: go run -race main.go

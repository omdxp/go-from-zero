package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(2 * gr)

	var n int = 0

	// step 1
	var m sync.Mutex // declare a mutex

	for i := 0; i < gr; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Microsecond)
			// step 2
			m.Lock() // blocks the access of the variable from other goroutines
			n++
			// step 3
			m.Unlock() // open the access again to n variable
		}()
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Microsecond)
			m.Lock()
			defer m.Unlock()
			n--
		}()
	}

	wg.Wait()
	fmt.Println(n) // 0 (there is no data race!)
}

// mutexes solve the problem of data races
// mutex stands for mutual exclusion
// explicit synchronization where variable's access is protected through synchronization primitive such as mutex
// channels also are useful for this kind of problem

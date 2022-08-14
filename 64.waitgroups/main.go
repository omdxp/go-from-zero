package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // to indicate that this goroutine has finished executing
	fmt.Println("f1 (foroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
}

func f2() {
	fmt.Println("f2 execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	var wg sync.WaitGroup // waitgroup

	wg.Add(1)

	go f1(&wg)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())

	f2()

	wg.Wait() // main goroutine waits for all other goroutines defined in the waitgroup (f1 in this case)
	fmt.Println("main execution stopped")
}

// waitgroups allow to block (wait) until goroutines within that waitgroup have been executed

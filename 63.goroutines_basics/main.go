package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (foroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
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
	fmt.Println("main execution started")
	fmt.Println("number of CPUs:", runtime.NumCPU())             // number of logical CPUs
	fmt.Println("number of goroutines:", runtime.NumGoroutine()) // number of goroutines that currently exists (now it is 1 because main is a goroutine)

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // max procs

	go f1()                                                      // starting a goroutine
	fmt.Println("number of goroutines:", runtime.NumGoroutine()) // 2 (main with f1)

	f2()

	time.Sleep(100 * time.Microsecond) // sleep to see f1 execution
	// f1 had enough time to complete its execution
	fmt.Println("main execution stopped")
}

// goroutines they are not threads, they're way cheaper, in fact their call stack are around 1-2 Kb where threads can be around 1-2 Mb
// Go supports concurrency which is NOT parallelism, meaning it deals with multiple processes at once

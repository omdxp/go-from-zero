package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	fmt.Println("This is foorbar()")
}

func main() {
	defer foo()
	bar()

	fmt.Println("Just a string after defering foo() and calling bar()")

	defer foobar()

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// code that works (read/write) with the file
}

// defer will execute the statement at the end of function
// if there were more that one defer statements, they'll be executed in the reverse order they've been defered
// behind the scenes, defered function calls are push into a stack (LIFO order)

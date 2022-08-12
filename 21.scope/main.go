package main

import (
	"fmt" // compile time warning, "fmt" is imported twice (other one with an alias)
	f "fmt"
	// "fmt" // compile time error, cannot import the package twice
)

const done = false // package scoped

func main() { // package scoped
	fmt.Println()
	f.Println(done)
	var x = 1 // local (block) scoped
	{
		var x = 2
		_ = x
	}
	fmt.Println(x)
}

// import statements are file scoped (are only visible in this file)
// variables or constants decalred outside functions are package scoped (visible to all files with the same package)
// variables declared between curly braces are block scoped (only visible in the block they're declared in)

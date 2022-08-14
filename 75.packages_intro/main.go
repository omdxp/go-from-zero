package main

import (
	"fmt"

	"github.com/Omar-Belghaouti/go-from-zero/75.packages_intro/greet"
	"github.com/Omar-Belghaouti/go-from-zero/75.packages_intro/numbers"
)

func main() {
	fmt.Println(numbers.Even(420)) // Even function is accessible from this package because it starts with uppercase letter
	fmt.Println(numbers.IsPrime(13))
	fmt.Println(numbers.Odd(69))
	greet.Greet()
}

// there are 2 types of packages in Go:
//	1. executable packages, where the program executes (main package)
//	2. non-executable packages, they're called libraries or dependencies where they're needed to the executable program (like fmt)
// there are packages that are relative directly to Go workspace src directory ($GOPATH/src) that can be imported simply by its name (import "x")
// to import local packages you need to use the name you have chose in go.mod file to import it (same as numbers package here)
// go mod init github.com/Omar-Belghaouti/go-from-zero was launched at the creation of this repo and every local package will be related to it when importing

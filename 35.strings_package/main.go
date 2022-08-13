package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println

	result := strings.Contains("I love Golang!", "love") // check if substring is contained in a string
	p(result)

	result = strings.ContainsAny("success", "xc") // check if characters are within a string (either "x" or "c" are within)
	p(result)

	result = strings.ContainsRune("Golang", 'g') // check if a rune is within a string
	p(result)

	n := strings.Count("cheese", "e") // counts the number of occurrences of a substring in a string
	p(n)

	n = strings.Count("Five", "") // it will return the number of runes + 1
	p(n)

	// ToUpper and ToLower functions returns copies (do not mutate)
	p(strings.ToLower("GO PyTHon TyPeScript"))
	p(strings.ToUpper("GO PyTHon TyPeScript"))

	p("go" == "go")
	p("GO" == "go")
	p(strings.ToLower("GO") == strings.ToLower("go"))
	p(strings.EqualFold("GO", "go")) // recommended way to compare strings when case sensitivity is not important

	myStr := strings.Repeat("ab", 10) // Repeat returns a new string repeated in a specific number of times
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2) // replace first 2 "." characters with ":"
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // replace all "." characters with ":"
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":") // ReplaceAll returns a new string replacing all old strings by a new one
	p(myStr)

	s := strings.Split("a,b,c", ",") // returns a slice of strings from a specified separator ("," in this case)
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = strings.Split("Go go GO!", "") // separate by each rune
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "-") // Join joins a slice of strings by a specified separator ("-" in this case)
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) // return a slice of strings separated by whitespaces (including special characters)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux! \n") // TrimSpace will remove all leading and trailing whitespaces
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello Gophers!!!!?", ".!?") // Trim with remove all leading and trailing characters specified in the cutset argument (".!?" in this case)
	fmt.Printf("%q\n", s2)
}

// ToLower and ToUpper will iterate to each rune
// EqualFold is used for case insensitivity

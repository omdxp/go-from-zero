package main

import "fmt"

func main() {
	// declare a variable
	var age int = 24
	fmt.Println("age is", age)

	// type inference
	var name = "Omar"
	fmt.Println("name is", name)

	// mute unused local variable warning
	_ = name

	// short declaration operator (works only in block scope)
	s := "Learing Go!"
	fmt.Println(s)

	// multiple declaration
	car, cost := "Ferrari", 42069 // tuple assignment
	fmt.Println(car, cost)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	j, i = i, j // swap variables
	fmt.Println(i, j)
}

// in Go, multiple statement can be written in one line seperated with semicolons ";"
// it is recommened if you know the initial value to use the short declaration operator ":="
// otherwise use the var keyword

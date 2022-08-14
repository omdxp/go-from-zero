package main

import "fmt"

// func sayHello(name string) { // compile time error, redeclared
// 	fmt.Printf("Hello %s!\n", name)
// }

func main() {
	fmt.Println("Scope means name visibility")
	sayHello("Omar")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water boiling point in fehrenheit:", tf)
}

// imports are file scoped and they should be imported wherever they are being used (even in the same package)
// functions, types, constants, variables are package scoped

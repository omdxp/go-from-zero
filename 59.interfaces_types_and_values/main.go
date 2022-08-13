package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape // abstract type value (zero value is nil)
	fmt.Printf("%T\n", s)

	ball := circle{radius: 2.5} // concrete value of type circle
	s = ball                    // s took the form of a circle (at runtime)

	print(s) // underlying type is circle so it's gonna print the circle info
	fmt.Printf("Type of s: %T\n", s)

	room := rectangle{width: 2, height: 3}
	s = room // now s took the form of a rectangle (at runtime)
	fmt.Printf("Type of s: %T\n", s)
}

// under the hood interface values can be thought of a pair of concrete value and dynamic type
// interface value holds a value of a specific underlying concrete type
// calling a method on an interface value executes the method of the same name on its underlying type
// this is what polymorphism is about, the interface type can take poly (many dynamic forms) at runtime
// variables have types known at compile time and can never change
// interfaces have dynamic types that may change during runtime

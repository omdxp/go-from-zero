package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64 // a method signature
	perimeter() float64
}

type rectangle struct { // implements shape interface because it imelement all the method signatures for shape interface
	width, height float64
}

type circle struct { // also implements shape interface for the same reason
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// func printCircle(c circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Perimeter:", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Perimeter:", r.perimeter())
// }

func printShape(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3., height: 2.1}

	// printCircle(c1)
	// printRectangle(r1)

	// polymorphism in action
	printShape(c1) // if circle does not implement one of the shape methods (or all), compile time error will occur
	printShape(r1)
}

// interface is a collection of method signatures that an object which is most of the time is a named type can implement
// interfaces define the behavior of an object and can acheive polymorphism
// interfaces are not generic types like they are in other languages like Java
// in Go interfaces are implemented implicitly
// "if it walks like duck, swings like a duck and quacks like a duck then it's a duck"
// a type can implement one or more interfaces

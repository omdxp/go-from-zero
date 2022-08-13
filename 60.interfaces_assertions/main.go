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
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	// s.volume() // compile time error, even though s has circle type (its underlying type)
	// s can only access only the methods that are defined inside the interface (area() and perimeter())

	ball, ok := s.(circle) // type assertion
	if ok {
		fmt.Printf("Ball volume: %v\n", ball.volume())
	}

	// type switches
	s = rectangle{width: 3.4, height: 2.2}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}
}

// in order to be able to invoke a method of the underlying dynamic type on the interface value we need to do "type assertions" or "type switches"

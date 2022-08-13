package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

// geometry is embedding shape and object interfaces
type geometry interface {
	shape  // all shape methods are added
	object // all object methods are added
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * math.Pow(c.edge, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c) // c must implements geometry interface
	fmt.Printf("Area: %v, Volume: %v\n", a, v)
}

// an interface cannot implement other interfaces or extends them, inheritance is not supported
// to acheive that, create a new interface by merging two or more interfaces "interface embedding"
// circular embedding of interfaces is not allowed and it is detected at compile time

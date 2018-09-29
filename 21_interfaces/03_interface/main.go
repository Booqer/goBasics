package main

import (
	"fmt"
	"math"
)

// create a user defined type of square
// the underlying type is a struct
type square struct {
	side float64
}

// create a user defined type of circle
// the underlying type is a struct
type circle struct {
	radius float64
}

// create a user defined type of shape
// the underlying type is interface
type shape interface {
	// area() float64 is the method signature
	// anything with that method signature implements the shape interface
	area() float64
}

// attach a method to the square type
// note how this method meets the method signature of the shape interface, thus can implement shape
func (s square) area() float64 {
	return s.side * s.side
}

// attach a method to the circle type
// note how this method meets the method signature of the shape interface, thus can implement shape
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// since circle and square both implement the shape interface anywhere that asks for a shape can take either a circle or a square
func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}

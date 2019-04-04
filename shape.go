package main

import "math"

// Shape is an interface that impelements Area and Perimeter
// common shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Square is the mathematical representation of a square
type Square struct {
	side float64
}

// Perimeter returns the perimeter of a square
func (s Square) Perimeter() float64 {
	return 4 * s.side
}

// Area returns the area of a square
func (s Square) Area() float64 {
	return s.side * s.side
}

// Rectangle is the mathematical representation of a rectangle
type Rectangle struct {
	length float64
	height float64
}

// Perimeter returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.height)
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.height
}

// Circle is the mathematical representation of a circle
type Circle struct {
	radius float64
}

// Perimeter returns the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

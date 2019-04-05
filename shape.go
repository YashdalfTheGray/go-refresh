package main

import (
	"fmt"
	"math"
)

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

// PrintShape prints out the right thing depending on shape
func PrintShape(s Shape) string {
	switch value := s.(type) {
	case Square:
		return fmt.Sprintf(
			"The square with side %.2f has perimeter %.2f and area %.2f.",
			value.side, value.Perimeter(), value.Area())
	case Rectangle:
		return fmt.Sprintf(
			"The rectangle with sides %.2f and %.2f has perimeter %.2f and area %.2f.",
			value.height, value.length, value.Perimeter(), value.Area())
	case Circle:
		return fmt.Sprintf(
			"The circle with radius %.2f has perimeter %.2f and area %.2f.",
			value.radius, value.Perimeter(), value.Area())
	default:
		return fmt.Sprintf(
			"The shape passed in has perimeter %.2f and area %.2f.",
			value.Perimeter(), value.Area())
	}
}

package main

import "fmt"

// Vertex2D represents a point in 2D space. Implements the
// fmt.Stringer interface
type Vertex2D struct {
	x, y float64
}

func (v Vertex2D) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", v.x, v.y)
}

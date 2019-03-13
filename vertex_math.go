package main

import "math"

// Distance returns the distance between two vertices in 2D space
func Distance(a, b Vertex2D) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}

// TranslateX translates a 2D vector in the horizontal direction
func TranslateX(v *Vertex2D, factor float64) {
	v.x *= factor
}

// TranslateY translates a 2D vector in the vertical direction
func TranslateY(v *Vertex2D, factor float64) {
	v.y *= factor
}

// Translate is a convenience wrapper around TranslateX and TranslateY
func Translate(v *Vertex2D, xFactor, yFactor float64) {
	TranslateX(v, xFactor)
	TranslateY(v, yFactor)
}

package main

type evaluate func(x, y int) int

// Picture provides a dx by dy picture with each element evaluated
// by the evaluation function passed in as fn.
func Picture(dx, dy int, fn evaluate) [][]uint8 {
	result := make([][]uint8, dy)

	for y := range result {
		result[y] = make([]uint8, dx)
		for x := range result[y] {
			result[y][x] = uint8(fn(x, y))
		}
	}

	return result
}

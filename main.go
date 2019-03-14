package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func main() {
	DetectOS()
	println("Running all sample code\n")

	SimpleHello()

	quotient, remainder := Divide(25, 4)
	fmt.Println(fmt.Sprintf("%d divided by %d is %d with a remainder of %d.", 25, 4, quotient, remainder))

	if result, err := FindOccurences("abc", 100, "b"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("There are %d %s's in the first %d letters of \"%s\" repeated.", result, "b", 100, "abc"))
	}

	a := Vertex2D{4, 5}
	b := Vertex2D{1, 1}
	initialDistance := Distance(a, b)
	fmt.Println(fmt.Sprintf("The distance between %s and %s is %.2f.", a, b, initialDistance))
	Translate(&a, 2, 4)
	updatedDistance := Distance(a, b)
	fmt.Println(fmt.Sprintf("The distance between %s and %s is %.2f.", a, b, updatedDistance))

	evaluationFn := func(x, y int) int {
		return x ^ y
	}
	wrapper := func(dx, dy int) [][]uint8 {
		return Picture(dx, dy, evaluationFn)
	}
	pic.Show(wrapper)
}

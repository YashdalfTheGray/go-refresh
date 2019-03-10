package main

import (
	"fmt"
)

func main() {
	SimpleHello()

	quotient, remainder := Divide(25, 4)
	fmt.Println(fmt.Sprintf("%d divided by %d is %d with a remainder of %d.", 25, 4, quotient, remainder))
}

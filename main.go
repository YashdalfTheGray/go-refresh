package main

import (
	"fmt"
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
}

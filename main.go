package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
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
	encoded, encodingError := GenerateAndShow(wrapper, false)
	if encodingError != nil {
		fmt.Println(encodingError)
	}
	if writeError := ioutil.WriteFile("bin/temp.base64", []byte(encoded), 0644); writeError != nil {
		fmt.Println(writeError)
	} else {
		fmt.Println("Wrote file bin/temp.base64 with image in base64 form")
	}

	s := "This is a string that we will use to detect word substring in"
	sub := "This string is substring"
	if result, subErr := SubstringWithMaps(s, sub); subErr != nil {
		fmt.Println(subErr)
	} else {
		fmt.Println(fmt.Sprintf("Does s containe all words in sub? %t", result))
	}

	fib := FibonacciGen()
	fmt.Println("the first ten terms of the Fibonacci sequence are")
	termsString := ""
	for i := 0; i < 10; i++ {
		termsString += strconv.Itoa(fib())
		if i < 9 {
			termsString += ", "
		}
	}
	fmt.Println(termsString)

	var mySquare, myRectangle, myCircle Shape
	mySquare = Square{side: 4}
	myRectangle = Rectangle{length: 3, height: 4}
	myCircle = Circle{radius: 5}
	fmt.Println(PrintShape(mySquare))
	fmt.Println(PrintShape(myRectangle))
	fmt.Println(PrintShape(myCircle))

	fmt.Println("Reading 32 b's from InfiniteLetters")
	bSlice := make([]byte, 32)
	il := InfiniteLetters{letter: 'b'}
	if n, ilReadErr := il.Read(bSlice); ilReadErr == nil {
		fmt.Println(string(bSlice))
		fmt.Println(fmt.Sprintf("Read %d characters from InfiniteLetters", n))
	} else {
		fmt.Println(ilReadErr)
	}

	fmt.Println("Reading \"this is a test message\", rot13 encoded, 8 bytes at a time")
	encodedSlice := make([]byte, 8)
	testMessage := Message("this is a test message")
	rot13Enc := testMessage.NewRot13Encoder()
	for {
		n, err := rot13Enc.Read(encodedSlice)
		fmt.Println(fmt.Sprintf("Read %d bytes with error %e", n, err))
		if err == io.EOF {
			break
		}
		fmt.Println("\"" + string(encodedSlice) + "\"")
	}

	fmt.Print("\n")
}

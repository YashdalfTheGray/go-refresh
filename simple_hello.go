package main

import "fmt"

// SimpleHello just prints out a simple hello world interpolated string
func SimpleHello() {
	fmt.Println(fmt.Sprintf("Hello, %s!\n", "World"))
}

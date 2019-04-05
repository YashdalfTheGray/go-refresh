package main

import "fmt"

// ArgError is a custom error class to tell the user about an
// invalid argument passed to a function. Implements the error
// interface
type ArgError struct {
	name    string
	message string
}

func (e ArgError) Error() string {
	return fmt.Sprintf("%s is invalid - %s", e.name, e.message)
}

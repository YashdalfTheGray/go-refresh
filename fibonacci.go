package main

// FibonacciGen returns a function that returns the next
// fibonacci number on every invocation
func FibonacciGen() func() int {
	termNumber := 1
	nMinusTwoTerm := 1
	nMinusOneTerm := 1

	return func() int {
		if termNumber < 3 {
			termNumber++
			return 1
		}
		currentTerm := nMinusOneTerm + nMinusTwoTerm
		nMinusTwoTerm = nMinusOneTerm
		nMinusOneTerm = currentTerm

		return currentTerm
	}
}

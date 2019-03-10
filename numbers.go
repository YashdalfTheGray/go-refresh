package main

// Divide divides two numbers and returns the quotient and remainder
func Divide(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor

	return quotient, remainder
}

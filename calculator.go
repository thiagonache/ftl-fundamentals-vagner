// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the first
// from the second.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply take two numbers and returns the result of multiplying
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide take two numbers and return result of division from the first
// from the second
func Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Can't divide by zero, the default value for expression is 0")
	}
	return a / b, nil
}

// SquareRoot take one number and return a number multiply by the informed number
func SquareRoot(a float64) (float64, error) {
	if a <= 0 {
		return 0, errors.New("Zero, or negative values is not possible")
	}
	return math.Sqrt(a), nil
}

package calc

import "errors"

// Add two integer numbers
func Add(x, y int) (int) {
	return x + y
}

// Subtract two integer numbers
func Subtract(x, y int) (int) {
	return x - y
}

// Multiply two integer numbers
func Multiply(x, y int) (int) {
	return x * y
}

// Divide two integer numbers
func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("zero division error")
	}

	return x / y, nil
}
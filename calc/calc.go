package calc

import "errors"

// Add two integer numbers
func Add(x, y int32) (int32) {
	return x + y
}

// Subtract two integer numbers
func Subtract(x, y int32) (int32) {
	return x - y
}

// Multiply two integer numbers
func Multiply(x, y int32) (int32) {
	return x * y
}

// Divide two integer numbers
func Divide(x, y int32) (int32, error) {
	if y == 0 {
		return 0, errors.New("zero division error")
	}

	return x / y, nil
}
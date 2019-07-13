package calc

import (
	"testing"
)

// Add two integer numbers
func TestAdd(t *testing.T) {
	result := Add(1, 1)
	expected := 2

	if result != expected {
		t.Errorf("Add(1, 1) = %d; expecred %d", result, expected)
	}
}

// Subtract two integer numbers
func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	expected := 3

	if result != expected {
		t.Errorf("Subtract(5, 2) = %d; expecred %d", result, expected)
	}
}

// Multiply two integer numbers
func TestMultiply(t *testing.T) {
	result := Multiply(3, 6)
	expected := 18

	if result != expected {
		t.Errorf("Multiply(3, 6) = %d; expecred %d", result, expected)
	}
}

// Divide two integer numbers
func TestDivide(t *testing.T) {
	result, err := Divide(20, 5)
	expected := 4

	if err != nil {
		t.Errorf("Divide(20, 5) returned error: %v", err)
	}

	if result != expected {
		t.Errorf("Divide(20, 5) = %d; expecred %d", result, expected)
	}
}

// Divide number by zero
func TestDivideByZero(t *testing.T) {
	_, err := Divide(20, 0)

	if err == nil {
		t.Errorf("Divide(20, 0) should return an error")
	}
}
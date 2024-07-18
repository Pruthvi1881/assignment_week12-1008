package main

import "testing"

func TestFactorial(t *testing.T) {
	// Test with a positive integer
	if result := Factorial(5); result != 120 {
		t.Errorf("Factorial(5) = %d; want 120", result)
	}

	// Test with zero
	if result := Factorial(0); result != 1 {
		t.Errorf("Factorial(0) = %d; want 1", result)
	}

	// Test with a negative integer
	if result := Factorial(-5); result != -1 {
		t.Errorf("Factorial(-5) = %d; want -1", result)
	}

	// Test with a large number
	if result := Factorial(10); result != 3628800 {
		t.Errorf("Factorial(10) = %d; want 3628800", result)
	}
}

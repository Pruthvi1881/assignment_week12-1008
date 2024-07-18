package main

import "fmt"

// Factorial calculates the factorial of a given non-negative integer.
// It returns 1 for 0 as 0! is 1.
func Factorial(n int) int {
	if n < 0 {
		return -1 // Return -1 for negative inputs to indicate an error
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(Factorial(5)) // Output: 120
}

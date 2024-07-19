package main

import "fmt"

// LCS finds the longest common subsequence of two strings.
func LCS(X, Y string) string {
	m := len(X)
	n := len(Y)

	// Create a 2D table to store lengths of longest common subsequence.
	L := make([][]int, m+1)
	for i := range L {
		L[i] = make([]int, n+1)
	}

	// Fill L[][] in bottom-up fashion.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = max(L[i-1][j], L[i][j-1])
			}
		}
	}

	// Reconstruct LCS string from the table.
	index := L[m][n]
	lcs := make([]byte, index)
	i, j := m, n
	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			lcs[index-1] = X[i-1]
			i--
			j--
			index--
		} else if L[i-1][j] > L[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}

// Helper function to find the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage of the LCS function.
	X := "ABCBDAB"
	Y := "BDCAB"
	fmt.Println("LCS of", X, "and", Y, "is", LCS(X, Y))
}

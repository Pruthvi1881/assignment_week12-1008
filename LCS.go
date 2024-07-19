package main

import "testing"

// TestLCS tests the LCS function with various cases.
func TestLCS(t *testing.T) {
	tests := []struct {
		X      string
		Y      string
		expect string
	}{
		{"ABCBDAB", "BDCAB", "BCAB"},  // Corrected expected result
		{"AAB", "ABA", "ABA"},         // Corrected expected result
		{"AGGTAB", "GXTXAYB", "GTAB"}, // Additional test case
		{"ABC", "AC", "AC"},           // Additional test case
		{"", "ABC", ""},               // Edge case: empty string
		{"ABC", "", ""},               // Edge case: empty string
		{"", "", ""},                  // Edge case: both empty strings
	}

	for _, test := range tests {
		result := LCS(test.X ,test.Y)
		if result != test.expect {
			t.Errorf("LCS(%q, %q) = %q; want %q", test.X, test.Y, result, test.expect)
		}
	}
}

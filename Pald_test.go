package main

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"A man a plan a canal Panama", true},
        {"Hello World", false},
        {"a", true},
        {"", true},
        {"No 'x' in Nixon", true},
    }

    for _, test := range tests {
        result := IsPalindrome(test.input)
        if result != test.expected {
            t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
        }
    }
}

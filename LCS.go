package main

import (
    "fmt"
    "strings"
    "unicode"
)

// IsPalindrome checks if a given string is a palindrome.
// It removes non-alphanumeric characters and ignores case.
func IsPalindrome(s string) bool {
    var cleaned strings.Builder
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsNumber(r) {
            cleaned.WriteRune(unicode.ToLower(r))
        }
    }
    // Get the cleaned string and its reverse
    cleanedStr := cleaned.String()
    reversedStr := reverse(cleanedStr)
    return cleanedStr == reversedStr
}

// Helper function to reverse a string
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    fmt.Println(IsPalindrome("A man a plan a canal Panama")) // Output: true
    fmt.Println(IsPalindrome("Hello World"))                  // Output: false
}

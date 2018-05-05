package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func reverseWords(s string) string {
	var result []string
	for _, w := range strings.Fields(s) {
		result = append(result, reverseString(w))
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(reverseWords("abc cba"))
}
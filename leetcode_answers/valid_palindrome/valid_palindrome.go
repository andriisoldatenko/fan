package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {}

func isPalindrome(s string) bool {
	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return unicode.ToLower(r)
	}, s)
	fmt.Println(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

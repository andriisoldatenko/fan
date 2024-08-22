package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x int
		n bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}
	for _, test := range tests {
		result := isPalindrome(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.n)
		}
	}
}

package main

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		x string
		n bool
	}{
		{"USA", true},
		{"FlaG", false},
		{"Google", true},
		{"leetcode", true},
		{"g", true},
		{"G", true},
	}
	for _, test := range tests {
		result := detectCapitalUse(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %v, want: %v.: func(%v)", result, test.n, test.x)
		}
	}
}

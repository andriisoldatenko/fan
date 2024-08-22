package main

import "testing"

func TestAddDigits(t *testing.T) {
	tests := []struct {
		x int
		n int
	}{
		{38, 2},
		{111, 3},
		{0, 0},
		{1, 1},
	}
	for _, test := range tests {
		result := addDigits(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %d, want: %d.", result, test.n)
		}
	}
}

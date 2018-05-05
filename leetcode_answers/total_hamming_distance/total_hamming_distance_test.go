package main

import "testing"


func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x int
		y int
		n int
	}{
		{1, 4, 2},
	}
	for _, test := range tests {
		total := hammingDistance(test.x, test.y)
		if total != test.n {
			t.Errorf("Total was incorrect, got: %d, want: %d.", total, test.n)
		}
	}
}
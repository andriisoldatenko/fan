package main

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		x string
		y string
		n int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, test := range tests {
		total := numJewelsInStones(test.x, test.y)
		if total != test.n {
			t.Errorf("Total was incorrect, got: %d, want: %d.", total, test.n)
		}
	}
}

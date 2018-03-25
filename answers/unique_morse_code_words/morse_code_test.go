package main

import "testing"


func TestUniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		x []string
		n int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
		{[]string{}, 0},
		{[]string{"gin"}, 1},
	}
	for _, test := range tests {
		total := uniqueMorseRepresentations(test.x)
		if total != test.n {
			t.Errorf("Total was incorrect, got: %d, want: %d.", total, test.n)
		}
	}
}
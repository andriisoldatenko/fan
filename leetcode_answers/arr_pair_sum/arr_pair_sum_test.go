package main

import "testing"

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		x []int
		n int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{1, 2}, 1},
		{[]int{7, 8, 6, 3}, 10},
		{[]int{7, 3, 1, 0, 0, 6}, 7},
	}
	for _, test := range tests {
		result := arrayPairSum(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %d, want: %d.", result, test.n)
		}
	}
}

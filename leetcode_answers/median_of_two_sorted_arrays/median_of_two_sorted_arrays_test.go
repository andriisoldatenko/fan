package main

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		x []int
		y []int
		m float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, test := range tests {
		result := findMedianSortedArrays(test.x, test.y)
		if result != test.m {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.m)
		}
	}
}

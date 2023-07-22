package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		x []int
		n int
		r []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	}
	for _, test := range tests {
		result := kidsWithCandies(test.x, test.n)
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.n)
		}
	}
}

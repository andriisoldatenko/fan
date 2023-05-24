package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums    []int
		target  int
		indices []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, test := range tests {
		indices := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(indices, test.indices) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", indices, test.indices)
		}
	}
}
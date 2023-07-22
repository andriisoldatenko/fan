package main

import (
	"github.com/google/go-cmp/cmp"
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
		got := twoSum(test.nums, test.target)
		if diff := cmp.Diff(test.indices, got); diff != "" {
			t.Errorf("twoSum() mismatch (-want +got):\n%s", diff)
		}

	}
}

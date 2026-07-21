package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		out  int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
	}
	for _, test := range tests {
		result := removeElement(test.nums, test.val)
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.out)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestLeftRightDifference(t *testing.T) {
	tests := []struct {
		x []int
		r []int
	}{
		{[]int{1}, []int{0}},
		{[]int{10, 4, 8, 3}, []int{15, 1, 11, 22}},
	}
	for _, test := range tests {
		result := leftRightDifference(test.x)
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.r)
		}
	}
}

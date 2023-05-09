package main

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		x []int
		y []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		result := runningSum(test.x)
		if !reflect.DeepEqual(result, test.y) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.x)
		}
	}
}

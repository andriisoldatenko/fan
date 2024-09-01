package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, test := range tests {
		result := productExceptSelf(test.in)
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.out)
		}
	}
}

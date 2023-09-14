package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		// nums []int
		// i    []int
	}{
		// {[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
	}
	for _, test := range tests {
		result := example(/*test.nums*/)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}

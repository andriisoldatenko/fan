package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		answer  string
	}{
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
	}
	for _, test := range tests {
		result := convert(test.s, test.numRows)
		if !reflect.DeepEqual(result, test.answer) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.answer)
		}
	}
}

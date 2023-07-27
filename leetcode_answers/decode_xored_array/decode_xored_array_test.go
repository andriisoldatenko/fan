package main

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	tests := []struct {
		nums   []int
		first  int
		result []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 0, 2, 1}},
	}
	for _, test := range tests {
		result := decode(test.nums, test.first)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.result)
		}
	}
}

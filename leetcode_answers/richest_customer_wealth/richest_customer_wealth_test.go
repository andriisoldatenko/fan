package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		nums [][]int
		i    int
	}{
		{[][]int{{1, 2, 3}, {3, 2, 1}}, 6},
	}
	for _, test := range tests {
		result := maximumWealth(test.nums)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	tests := []struct {
		x int
		y int
		z []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{1, 1, []int{1}},
		{99, 99, []int{99}},
	}
	for _, test := range tests {
		result := selfDividingNumbers(test.x, test.y)
		if !reflect.DeepEqual(result, test.z) {
			t.Errorf("Total was incorrect, got: %b, want: %b.", result, test.z)
		}
	}
}

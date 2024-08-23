package main

import (
	"reflect"
	"testing"
)

func Test_findComplement(t *testing.T) {
	tests := []struct {
		num int
		i   int
	}{
		{5, 2},
		{1, 0},
		{0, 1},
	}
	for _, test := range tests {
		result := findComplement(test.num)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}

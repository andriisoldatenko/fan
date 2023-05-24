package main

import (
	"reflect"
	"testing"
)


func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		x []int
		n []int
	}{
		{[]int{1, 2}, []int{ 1, 2, 1, 2 }},
	}
	for _, test := range tests {
		result := getConcatenation(test.x)
		if !reflect.DeepEqual(result, test.n) {
			t.Errorf("Total was incorrect, got: %v, want: %v.: func(%v)", result, test.n, test.x)
		}
	}
}

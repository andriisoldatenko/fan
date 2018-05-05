package main

import "testing"


func TestReverse(t *testing.T) {
	tests := []struct {
		x int
		n int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
		{-2147483648, 0},
	}
	for _, test := range tests {
		result := reverse(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.n)
		}
	}
}
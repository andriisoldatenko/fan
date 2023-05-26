package main

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		s string
		i int
	}{
		{"a", 1},
		{"abccccdd", 7},
		{"aaa", 3},
		{"aa", 2},
	}
	for _, test := range tests {
		result := longestPalindrome(test.s)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}

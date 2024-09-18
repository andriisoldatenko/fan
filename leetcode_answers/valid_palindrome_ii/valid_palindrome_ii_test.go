package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		s string
		r bool
	}{
		{"a", true},
		{"ab", true},
		{"aba", true},
		{"abca", true},
		{"abc", false},
		{"eedede", true},
	}
	for _, test := range tests {
		result := validPalindrome(test.s)
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.r)
		}
	}
}

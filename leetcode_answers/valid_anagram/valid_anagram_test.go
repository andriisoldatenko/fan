package main

import (
	"reflect"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		s  string
		s1 string
		r  bool
	}{
		{"", "", true},
		{"rat", "car", false},
		{"anagram", "nagaram", true},
	}
	for _, test := range tests {
		result := isAnagram(test.s, test.s1)
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.r)
		}
	}
}

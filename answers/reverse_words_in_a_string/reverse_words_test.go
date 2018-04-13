package main

import "testing"


func TestReverseString(t *testing.T) {
	tests := []struct {
		x string
		n string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"abc", "cba"},
		{"lazy 犬", "yzal 犬"},
	}
	for _, test := range tests {
		result := reverseWords(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.n)
		}
	}
}
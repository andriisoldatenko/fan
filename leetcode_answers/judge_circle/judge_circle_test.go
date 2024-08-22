package main

import "testing"

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		x string
		n bool
	}{
		{"UD", true},
		{"LL", false},
	}
	for _, test := range tests {
		result := judgeCircle(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.n)
		}
	}
}

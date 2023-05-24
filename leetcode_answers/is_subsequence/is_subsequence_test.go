package main

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"abc", "ahbgdc"}, true},
		{"test2", args{"axc", "ahbgdc"}, false},
		{"test3", args{"", "ahbgdc"}, true},
		{"test31", args{"abc", ""}, false},
		{"test4", args{"b", "abc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

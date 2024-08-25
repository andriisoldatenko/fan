package main

import (
	"reflect"
	"testing"
)

func Test_scoreOfString(t *testing.T) {
	tests := []struct {
		s string
		i int
	}{
		{"hello", 13},
		{"zaz", 50},
	}
	for _, test := range tests {
		result := scoreOfString(test.s)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("DeepEqual was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}

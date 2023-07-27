package main

import (
	"reflect"
	"testing"
)

func Test_mostWordsFound(t *testing.T) {
	tests := []struct {
		s []string
		i int
	}{
		{[]string{"a"}, 1},
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
	}
	for _, test := range tests {
		result := mostWordsFound(test.s)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		names   []string
		heights []int
		result  []string
	}{
		{[]string{"Mary", "John", "Emma"}, []int{180, 165, 170}, []string{"Mary", "Emma", "John"}},
	}
	for _, test := range tests {
		result := sortPeople(test.names, test.heights /*test.nums*/)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.result)
		}
	}
}

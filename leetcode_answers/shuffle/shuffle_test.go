
package main

import (
	"testing"
	"reflect"
)


func TestShuffle(t *testing.T) {
	tests := []struct {
		nums []int
		n int
		r []int
	}{
		{ []int{2,5,1,3,4,7}, 3, []int{2,3,5,4,1,7} },
	}
	for _, test := range tests {
		result := shuffle(test.nums, test.n)
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %b, want: %b.", result, test.r)
		}
	}
}

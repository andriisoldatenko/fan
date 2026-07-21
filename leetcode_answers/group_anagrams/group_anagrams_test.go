package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		strs []string
		r    [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}}},
	}
	for _, test := range tests {
		result := groupAnagrams(test.strs)

		sort.SliceStable(result[0])

		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.r)
		}
	}
}

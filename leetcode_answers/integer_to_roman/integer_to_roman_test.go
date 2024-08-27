package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		i int
		r string
	}{
		{5, "V"},
		{15, "XV"},
		{2005, "MMV"},
		{4, "IV"},
		{1994, "MCMXCIV"},
		{58, "LVIII"},
		{3749, "MMMDCCXLIX"},
	}
	for _, test := range tests {
		result := intToRoman(test.i)
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.r)
		}
	}
}

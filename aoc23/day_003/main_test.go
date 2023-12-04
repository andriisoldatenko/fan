package main

import (
	"reflect"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected []Number
	}{
		{"11..", []Number{{start: 0, end: 1, row: 0, value: 11}}},
		{"11..11", []Number{{start: 0, end: 1, row: 0, value: 11}, {start: 4, end: 5, row: 0, value: 11}}},
		{"..35..633.", []Number{{start: 2, end: 3, row: 0, value: 35}, {start: 6, end: 8, row: 0, value: 633}}},
	}

	for _, tt := range tests {
		actual := findNumbers(tt.input, 0)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Fatalf("expected=%v, got=%v",
				tt.expected, actual)
		}
	}
}

///.....
// .11..
// .....
///

func Test_hasNeighbours(t *testing.T) {
	board := [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '1', '1', '.', '.'},
		{'.', '.', '.', '.', '.'},
	}
	actual := includeNumber(1, 2, 1, board)
	if !reflect.DeepEqual(actual, false) {
		t.Fatalf("expected=%v, got=%v",
			false, actual)
	}

	board = [][]rune{
		{'.', '.', '.', '.', '.'},
		{'.', '1', '1', '.', '.'},
		{'.', '.', '#', '.', '.'},
	}
	actual = includeNumber(1, 2, 1, board)
	if !reflect.DeepEqual(actual, true) {
		t.Fatalf("expected=%v, got=%v",
			true, actual)
	}
}

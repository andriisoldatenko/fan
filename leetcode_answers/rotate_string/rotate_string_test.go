package main

import "testing"

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		x string
		n string
	}{
		{"hello", "olleh"},
		{"abc", "cba"},
		{"The quick brown 狐 jumped over the lazy 犬", "犬 yzal eht revo depmuj 狐 nworb kciuq ehT"},
	}
	for _, test := range tests {
		result := reverseString(test.x)
		if result != test.n {
			t.Errorf("Total was incorrect, got: %s, want: %s.", result, test.n)
		}
	}
}

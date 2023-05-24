package main

func isSubsequence(s string, t string) bool {
	// Input: s = "abc", t = "ahbgdc"
	i, j := 0, 0
	count := 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i += 1
			count += 1
		}
		j += 1
	}
	if count == len(s) {
		return true
	}
	return false
}

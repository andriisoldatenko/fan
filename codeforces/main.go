package codeforces


func isPalindrome(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var j = len(s1) - 1
	for i := 0; i <= j; i++ {
		if s1[j-i] != s2[i] {
			return false
		}
	}
	return true
}
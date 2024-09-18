package main

func main() {}

func RemoveIndex(s string, index int) string {
	ret := make([]byte, 0)
	ret = append(ret, s[:index]...)
	return string(append(ret, s[index+1:]...))
}

func validPalindrome(s string) bool {
	isPalindrome := true

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		return isPalindrome
	}

	for j := 0; j < len(s); j++ {
		isPalindrome = true
		ns := RemoveIndex(s, j)
		for i := 0; i < len(ns)/2; i++ {
			if ns[i] != ns[len(ns)-i-1] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			return isPalindrome
		}
	}

	return isPalindrome
}

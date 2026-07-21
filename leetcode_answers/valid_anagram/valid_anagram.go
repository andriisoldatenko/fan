package main

func main() {}

// https://leetcode.com/problems/valid-anagram/
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	m := make(map[rune]int)
//	for _, r := range s {
//		m[r]++
//	}
//	for _, r := range t {
//		m[r]--
//	}
//
//	for _, v := range m {
//		if v != 0 {
//			return false
//		}
//	}
//	return true
//}

// without utf-8 support but faster
func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		m[rune(s[i])]++
		m[rune(t[i])]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

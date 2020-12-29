package main

import "strings"

func duplicate_count(s1 string) int {
	m := make(map[string]int)
	for _, char := range s1 {
		normalized := strings.ToLower(string(char))
		_, ok := m[normalized]
		if ok {
			m[normalized]++
		} else {
			m[normalized] = 1
		}
	}
	total := 0
	for _, v := range m {
		if v > 1 {
			total += 1
		}
	}
	return total
}

func main() {}

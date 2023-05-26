package main

func longestPalindrome(s string) int {
	total := 0
	pairs := map[string]int{}
	for _, ch := range s {
		pairs[string(ch)] += 1
		if pairs[string(ch)]%2 == 0 {
			total += 2
			pairs[string(ch)] = 0
		}
	}

	max := 0
	for _, v := range pairs {
		if v > max {
			max = v
		}
	}
	return total + max
}

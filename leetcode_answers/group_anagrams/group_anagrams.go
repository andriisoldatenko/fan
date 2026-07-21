package main

import (
	"maps"
	"slices"
	"strconv"
)

func main() {}

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	mm := map[[26]int][]string{}
	for _, str := range strs {
		compressed := compress(str)
		mm[compressed] = append(mm[compressed], str)
	}

	res := [][]string{}
	for _, v := range mm {
		res = append(res, v)
	}

	return res
}

func compress(s string) [26]int {
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] += 1
	}

	return m
}

func compress2(s string) string {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	ss := ""

	for _, k := range slices.Sorted(maps.Keys(m)) {
		ss += string(k) + strconv.Itoa(m[k])
	}

	return ss
}

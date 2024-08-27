// https://leetcode.com/problems/integer-to-roman/description/
package main

import (
	"sort"
)

func main() {}

var mapping = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

var fournine = map[int]string{
	4:   "IV",
	9:   "IX",
	40:  "XL",
	90:  "XC",
	400: "CD",
	900: "CM",
}

func firstDigit(x int) int {
	for x > 9 {
		x /= 10
	}
	return x
}

func intToRoman(num int) string {
	result := ""

	for num != 0 {
		m := _intToRoman(num)
		first := firstDigit(num)
		if (first == 4) || (first == 9) {
			result += fournine[m]
		} else {
			result += mapping[m]
		}
		num = num - m
		m = num
	}

	return result
}

func _intToRoman(num int) int {
	first := firstDigit(num)

	if (first == 4) || (first == 9) {
		return maxSubstract(num, fournine)
	} else {
		return maxSubstract(num, mapping)
	}
}

func maxSubstract(num int, m map[int]string) int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	for _, i := range keys {
		if num >= i {
			return i
		}
	}
	return 0
}

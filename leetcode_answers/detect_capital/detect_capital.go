package main

import (
	"fmt"
)

func sum(sl []int) int {
	result := 0
	for _, i := range sl {
		result += i
	}
	return result
}
func detectCapitalUse(word string) bool {
	var caps []int
	for _, c := range word {
		if c <= 90 {
			caps = append(caps, 1)
		} else {
			caps = append(caps, 0)
		}
	}
	fmt.Println(sum(caps))
	if sum(caps) == 0 {
		return true
	}

	if len(caps) == sum(caps) || (caps[0] == 1 && sum(caps) == 1) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(detectCapitalUse("Google"))
}

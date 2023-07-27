package main

import "strings"

func main() {}

func mostWordsFound(sentences []string) int {
	maxWords := 0
	for _, s := range sentences {
		max := strings.Count(s, " ") + 1
		if max > maxWords {
			maxWords = max
		}
	}
	return maxWords
}

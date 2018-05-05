package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	morse_code := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
	                       "....","..",".---","-.-",".-..","--","-.","---",
	                       ".--.","--.-",".-.","...","-","..-","...-",".--",
	                       "-..-","-.--","--.."}
	// Index of first ascii letter
	a := []rune("a")[0]
	set := make(map[string]struct{})
	for _, word := range words {
		res := ""
		for _, w := range word {
			res += morse_code[w-a]
		}
		set[res] = struct{}{}
	}
	return len(set)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./review.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text() // GET the line string
		firstNumber := int32(0)
		for _, ch := range replace(t) {
			if unicode.IsDigit(ch) {
				firstNumber = ch - 48
				break
			}
		}
		secondNumber := int32(0)
		for _, ch := range ReverseReplace(t) {
			if unicode.IsDigit(ch) {
				secondNumber = ch - 48
				break
			}
		}
		fmt.Printf("%s %d\n", t, int(firstNumber)*10+int(secondNumber))
		result = result + int(firstNumber)*10 + int(secondNumber)
	}
	fmt.Println(result)
}

func replace(s string) string {
	args := []string{"three", "3", "eight", "8", "one", "1", "two", "2", "four", "4", "five", "5", "six", "6", "seven", "7", "nine", "9"}
	replacer := strings.NewReplacer(args...)
	return replacer.Replace(s)
}

func ReverseReplace(s string) string {
	args := []string{"eerht", "3", "thgie", "8", "eno", "1", "owt", "2", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "enin", "9"}
	replacer := strings.NewReplacer(args...)
	return replacer.Replace(Reverse(s))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

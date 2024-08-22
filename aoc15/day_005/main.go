package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	//part2()
}

func countVowels(s string) int {
	return strings.Count(s, "a") + strings.Count(s, "e") +
		strings.Count(s, "i") + strings.Count(s, "o") + strings.Count(s, "u")
}

func isTwice(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func part1() {
	disallowed := []string{"ab", "cd", "pq", "xy"}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		isNice := true
		for _, da := range disallowed {
			if strings.Contains(t, da) {
				isNice = false
				break
			}
		}
		if countVowels(t) < 3 {
			isNice = false
		}

		if !isTwice(t) {
			isNice = false
		}

		if isNice {
			result += 1
		}
	}
	fmt.Println(result)
}

func part2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		fmt.Println(t)
	}
	fmt.Println(result)
}

func strToints(str string) []int {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)
	result := make([]int, len(nums))
	for i, s := range nums {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func generatePalindrome2(min_, max_ string) []int {
	result := []int{}
	minL := len(min_)
	maxL := len(max_)
	for length := minL; length <= maxL; length++ {
		for _, xx := range generatePalindromeWithLength2(length) {
			_ = xx
			if xx <= atoi(max_) && xx >= atoi(min_) {
				result = append(result, xx)
			}
		}
	}

	return result
}

func part2() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_02/input_part2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	rr := []int{}
	for sc.Scan() {
		t := sc.Text() // GET the line string
		numbers := strings.Split(t, ",")
		for _, num := range numbers {
			start, end := strings.Split(num, "-")[0], strings.Split(num, "-")[1]
			fmt.Println(start, end)
			r := generatePalindrome2(start, end)
			fmt.Println(r)
			rr = append(rr, r...)
		}
	}
	fmt.Println(sum(rr))
}

func main() {
	//part1()
	part2()

	//fmt.Println(generatePalindromeWithLength2(2))
	//fmt.Println(generatePalindromeWithLength2(3))
	//fmt.Println(generatePalindromeWithLength2(5))
	//fmt.Println(generatePalindromeWithLength2(4))
	//fmt.Println(generatePalindromeWithLength2(6))
	//generatePalindromeWithLength2(10)
	//fmt.Println(sieveOfEratosthenes((10 / 2) + 1))
}

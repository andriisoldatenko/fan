package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// return list of primes less than N
func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func sum(arr []int) int {
	result := 0
	for _, a := range arr {
		result += a
	}
	return result
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func part1() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_02/input.txt")
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
			//fmt.Println(start, end)
			r := generatePalindrome(start, end)
			//fmt.Println(r)
			rr = append(rr, r...)
		}
	}
	fmt.Println(sum(rr))
}

func generatePalindrome(min_, max_ string) []int {
	result := []int{}
	minL := len(min_)
	maxL := len(max_)
	for length := minL; length <= maxL; length++ {
		for _, xx := range generatePalindromeWithLength(length) {
			_ = xx
			if xx <= atoi(max_) && xx >= atoi(min_) {
				result = append(result, xx)
			}
		}
	}

	return result
}

func atoi(num string) int {
	n, _ := strconv.Atoi(num)
	return n
}

func generatePalindromeWithLength(length int) []int {
	result := []int{}
	if length < 1 {
		return result
	}
	//
	//if length == 1 {
	//	for i := '0'; i <= '9'; i++ {
	//		result = append(result, int(i)-'0')
	//	}
	//	return result
	//}

	if length%2 != 0 {
		return result

	} else {
		halfLength := length / 2
		for j := powInt(10, halfLength-1); j < powInt(10, halfLength); j++ {
			n, _ := strconv.Atoi(fmt.Sprintf("%d%d", j, j))
			result = append(result, n)
		}
	}
	return result
}

func generatePalindromeWithLength2(length int) []int {
	result := []int{}

	primes := sieveOfEratosthenes(length + 1)
	for _, p := range primes {
		if length%p == 0 {
			l := length / p
			for j := powInt(10, l-1); j < powInt(10, l); j++ {
				n, _ := strconv.Atoi(strings.Repeat(fmt.Sprintf("%d", j), p))
				if !slices.Contains(result, n) {
					result = append(result, n)
				}
			}
		}
	}
	slices.Sort(result)
	return result
}

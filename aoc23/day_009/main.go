package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_009/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	total := []int{}
	for sc.Scan() {
		t := sc.Text()
		result := []int{}
		nums := strToints(t)
		fmt.Println(nums)

		last := nums[len(nums)-1]
		result = append(result, last)

		for {
			nums = findDiff(nums)
			fmt.Println(nums)
			result = append(result, nums[len(nums)-1])
			if allEqual(nums) {
				break
			}
		}
		//fmt.Println(result)
		tt := result[len(result)-1]
		for i := len(result) - 2; i >= 0; i-- {
			tt = tt + result[i]
		}
		fmt.Println(tt)
		total = append(total, tt)
	}
	fmt.Println(sum(total))

}

func part2() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_009/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	total := []int{}
	for sc.Scan() {
		t := sc.Text()
		result := []int{}
		nums := strToints(t)
		//fmt.Println(nums)

		last := nums[0]
		result = append(result, last)

		for {
			nums = findDiff(nums)
			//fmt.Println(nums)
			result = append(result, nums[0])
			if allEqual(nums) {
				break
			}
		}
		//fmt.Println(result)
		result = reverse(result)
		tt := result[0]
		for i := 1; i < len(result); i++ {
			tt = result[i] - tt
		}

		fmt.Println(tt)
		total = append(total, tt)
	}
	fmt.Println(sum(total))

}

func sum(in []int) int {
	r := 0
	for _, i := range in {
		r += i
	}
	return r
}

func allEqual(in []int) bool {
	first := 0
	for i := 1; i < len(in); i++ {
		if first != in[i] {
			return false
		}
	}
	return true
}
func findDiff(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr)-1; i++ {
		dfff := arr[i+1] - arr[i]
		result = append(result, dfff)
	}
	return result
}

func main() {
	part1()
	part2()
}

func strToints(str string) []int {
	nums := strings.Fields(str)
	result := make([]int, len(nums))
	for i, s := range nums {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

func reverse[S ~[]E, E any](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

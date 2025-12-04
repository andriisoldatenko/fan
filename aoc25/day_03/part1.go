package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func arrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")
}

func atoi(num string) int {
	n, _ := strconv.Atoi(num)
	return n
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

func findMax(ints []int) (int, int) {
	max1 := ints[0]
	index := 0

	for i := 0; i < len(ints); i++ {
		if ints[i] > max1 {
			max1 = ints[i]
			index = i
		}
	}
	return max1, index
}

func part1() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	result := []int{}
	ln := 1
	for sc.Scan() {
		t := sc.Text() // GET the line string

		ints := make([]int, len(t))

		for i, s := range t {
			ints[i] = atoi(string(s))
		}

		max1 := ints[0]
		max2 := -1
		max1Index := 0

		for i := 1; i < len(ints)-1; i++ {
			if ints[i] > max1 {
				max1 = ints[i]
				max1Index = i
			}
		}

		for i := max1Index + 1; i < len(ints); i++ {
			if ints[i] > max2 {
				max2 = ints[i]
			}
		}

		//fmt.Println(ln, ints)
		fmt.Println(atoi(fmt.Sprintf("%d%d", max1, max2)))
		ln++
		result = append(result, atoi(fmt.Sprintf("%d%d", max1, max2)))
	}
	fmt.Println(sum(result))
}

func part2() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	result := []int{}
	ln := 1
	for sc.Scan() {
		t := sc.Text() // GET the line string

		ints := make([]int, len(t))

		for i, s := range t {
			ints[i] = atoi(string(s))
		}
		maxs := make([]int, 12)
		startx := 0

		for mm := 0; mm < 12; mm++ {
			endx := len(ints) - (12 - mm) + 1
			chunk := ints[startx:endx]
			m, index := findMax(chunk)
			maxs[mm] = m
			startx = startx + index + 1
		}

		//fmt.Println(ln, ints)
		fmt.Println(arrayToString(maxs))
		ln++
		result = append(result, atoi(arrayToString(maxs)))
	}
	fmt.Println(sum(result))
}

func main() {
	//part1()
	part2()
}

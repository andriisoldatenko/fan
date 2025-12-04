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

func stringToSlice(s string) []string {
	return strings.Split(s, "")
}

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

func printArr(a [10][10]string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}

func isValidPos(i, j, n, m int) bool {
	if i < 0 || j < 0 || i >= n || j >= m {
		return false
	}
	return true
}

type Index struct {
	i int
	j int
}

func part1() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	result := 0
	a := [135][135]string{}
	i := 0
	for sc.Scan() {
		t := sc.Text() // GET the line string

		for j := 0; j < len(t); j++ {
			a[i][j] = string(t[j])
			fmt.Print(a[i][j])
		}
		fmt.Println()
		i++
	}
	//fmt.Println(a)
	// Size of a given 2d array
	n := len(a)
	m := len(a[0])

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {

			if a[i][j] == "@" {
				neighbours := []string{}
				indexes := []Index{}

				if isValidPos(i-1, j-1, n, m) {
					neighbours = append(neighbours, a[i-1][j-1])
					indexes = append(indexes, Index{i - 1, j - 1})
				}
				if isValidPos(i-1, j, n, m) {
					neighbours = append(neighbours, a[i-1][j])
				}
				if isValidPos(i-1, j+1, n, m) {
					neighbours = append(neighbours, a[i-1][j+1])
				}
				if isValidPos(i, j-1, n, m) {
					neighbours = append(neighbours, a[i][j-1])
				}
				if isValidPos(i, j+1, n, m) {
					neighbours = append(neighbours, a[i][j+1])
				}
				if isValidPos(i+1, j-1, n, m) {
					neighbours = append(neighbours, a[i+1][j-1])
				}
				if isValidPos(i+1, j, n, m) {
					neighbours = append(neighbours, a[i+1][j])
				}
				if isValidPos(i+1, j+1, n, m) {
					neighbours = append(neighbours, a[i+1][j+1])
				}
				fmt.Println(neighbours)

				cnt := strings.Count(strings.Join(neighbours, ""), "@")
				if cnt < 4 {
					result++
				}
			}
		}
	}
}

func part2() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	a := [135][135]string{}
	i := 0
	for sc.Scan() {
		t := sc.Text()
		for j := 0; j < len(t); j++ {
			a[i][j] = string(t[j])
		}
		i++
	}
	result := 0
	r := -1
	for r != 0 {
		rr, idx := process(a)
		fmt.Println(rr, idx)
		for _, v := range idx {
			a[v.i][v.j] = "x"
		}
		r = rr
		result += rr
	}
	fmt.Println(result)
}

func process(a [135][135]string) (int, []Index) {
	n := len(a)
	m := len(a[0])
	result := 0
	indexes := []Index{}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {

			if a[i][j] == "@" {
				neighbours := []string{}

				if isValidPos(i-1, j-1, n, m) {
					neighbours = append(neighbours, a[i-1][j-1])
				}
				if isValidPos(i-1, j, n, m) {
					neighbours = append(neighbours, a[i-1][j])
				}
				if isValidPos(i-1, j+1, n, m) {
					neighbours = append(neighbours, a[i-1][j+1])
				}
				if isValidPos(i, j-1, n, m) {
					neighbours = append(neighbours, a[i][j-1])
				}
				if isValidPos(i, j+1, n, m) {
					neighbours = append(neighbours, a[i][j+1])
				}
				if isValidPos(i+1, j-1, n, m) {
					neighbours = append(neighbours, a[i+1][j-1])
				}
				if isValidPos(i+1, j, n, m) {
					neighbours = append(neighbours, a[i+1][j])
				}
				if isValidPos(i+1, j+1, n, m) {
					neighbours = append(neighbours, a[i+1][j+1])
				}
				//fmt.Println(neighbours)

				cnt := strings.Count(strings.Join(neighbours, ""), "@")
				if cnt < 4 {
					indexes = append(indexes, Index{i, j})
					result++
				}
			}
		}

	}
	return result, indexes
}

func main() {
	//part1()
	part2()
}

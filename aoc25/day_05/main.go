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

func print_(obj any) {
	fmt.Println(obj)
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

type Range struct {
	x1 int64
	x2 int64
}

func in(a int64, r Range) bool {
	if a >= r.x1 && a <= r.x2 {
		return true
	}
	return false
}

func part1() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	rr := []Range{}
	nums := []int64{}

	for sc.Scan() {
		t := sc.Text()
		print_(t)

		if strings.Contains(t, "-") {
			// process ranges
			a := strings.Split(t, "-")
			rr = append(rr, Range{x1: int64(atoi(a[0])), x2: int64(atoi(a[1]))})
		} else if t != "" {
			nums = append(nums, int64(atoi(t)))
		}
	}

	fresh := map[int64]bool{}
	for _, n := range nums {
		for _, r := range rr {
			if in(n, r) {
				//print_(n)
				fresh[n] = true
			}
		}
	}
	print_("----")
	print_(len(fresh))

}

func part2() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	rr := []Range{}

	for sc.Scan() {
		t := sc.Text()

		if strings.Contains(t, "-") {
			// process ranges
			a := strings.Split(t, "-")
			rr = append(rr, Range{x1: int64(atoi(a[0])), x2: int64(atoi(a[1]))})
		}
	}

	for i := 0; i < len(rr); i++ {
		for j := i + 1; j < len(rr); j++ {
			// check overlap
			if in(rr[i].x1, rr[j]) || in(rr[i].x2, rr[j]) || in(rr[j].x1, rr[i]) || in(rr[j].x2, rr[i]) {
				// merge
				newRange := Range{
					x1: int64(math.Min(float64(rr[i].x1), float64(rr[j].x1))),
					x2: int64(math.Max(float64(rr[i].x2), float64(rr[j].x2))),
				}
				print_(fmt.Sprintf("%d %d", j, j+1))
				rr = append(rr[:j], rr[j+1:]...)
				print_(fmt.Sprintf("%d %d", i, i+1))
				rr = append(rr[:i], rr[i+1:]...)
				rr = append(rr, newRange)
				i = -1 // restart outer loop
				break
			}
		}
	}
	rrr := int64(0)
	for _, r := range rr {
		rrr += (r.x2 - r.x1) + 1
	}
	print_(rr)
	print_(rrr)
}

func main() {
	//part1()
	part2()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func stringToSlice(s string) []string {
	return strings.Split(s, "")
}

func arrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")
}

func atoi(num string) int64 {
	n, _ := strconv.Atoi(num)
	return int64(n)
}

func print_(obj any) {
	fmt.Println(obj)
}

func printM(arr [15][15]string) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
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

func mul(a [4]int64) int64 {
	result := int64(1)
	for _, v := range a {
		result *= v
	}
	return result
}

func sum_(a [4]int64) int64 {
	result := int64(0)
	for _, v := range a {
		result += v
	}
	return result
}

func mul_(a []int64) int64 {
	result := int64(1)
	for _, v := range a {
		result *= v
	}
	return result
}

func sum__(a []int64) int64 {
	result := int64(0)
	for _, v := range a {
		result += v
	}
	return result
}

func in(a int64, r Range) bool {
	if a >= r.x1 && a <= r.x2 {
		return true
	}
	return false
}

func part1() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	nums := []int64{}
	re := regexp.MustCompile("[0-9]+")
	//fmt.Println(re.FindAllString("abc123def987asdf", -1))
	A := [1000][4]int64{}
	i := 0
	result := int64(0)
	for sc.Scan() {
		t := sc.Text()
		//print_(t)

		_ = nums
		ns := re.FindAllString(t, -1)
		for j, n := range ns {
			A[j][i] = atoi(n)
		}
		if len(ns) == 0 {
			s := strings.Join(strings.Fields(t), "")
			fmt.Println(s)

			for ii, ch := range s {
				switch ch {
				case '*':
					result += mul(A[ii])
				case '+':
					result += sum_(A[ii])
				}
			}

		}

		i++
	}

	fmt.Println(result)

}

func part2() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_06/input_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")
	i := 0
	//result := int64(0)
	a := [5][3763]string{}
	for sc.Scan() {
		t := sc.Text()
		print_(t)
		for j, ch := range t {
			a[i][j] = string(ch)
		}
		i++
	}

	//printM(a)
	op := ""
	rr := int64(0)
	nums := []int64{} // change me
	for i := 0; i < 3763; i++ {
		s := ""
		for j := 0; j < 5; j++ {
			s += a[j][i]
		}
		ns := re.FindAllString(s, -1)

		if strings.Contains(s, "*") {
			op = "*"
		} else if strings.Contains(s, "+") {
			op = "+"
		}
		if len(ns) != 0 {
			nums = append(nums, atoi(ns[0]))
		}

		if ns == nil || i == (3763-1) {
			if op == "*" {
				rr += mul_(nums)
				print_(mul_(nums))
				nums = []int64{}
			}

			if op == "+" {
				rr += sum__(nums)
				print_(sum__(nums))
				nums = []int64{}
			}
			continue
		}
		//_ = re
		//fmt.Println(s)
		//fmt.Println()
	}
	print_(rr)

}

func main() {
	//part1()
	part2()
}

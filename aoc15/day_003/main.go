package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//part1()
	part2()
}

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		i := 0
		j := 0
		res := map[string]bool{"00": true}
		for _, ch := range t {
			if ch == '^' {
				j += 1
			}
			if ch == 'v' {
				j -= 1
			}
			if ch == '>' {
				i += 1
			}
			if ch == '<' {
				i -= 1
			}

			res[fmt.Sprintf("%d%d", i, j)] = true
		}
		result += len(res)

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
		t := sc.Text() // GET the line string
		x := 0
		y := 0
		xr := 0
		yr := 0
		res := map[string]bool{"0,0": true}
		for i, ch := range t {
			if ch == '^' {
				if i%2 == 0 {
					y += 1
				} else {
					yr += 1
				}
			}
			if ch == 'v' {
				if i%2 == 0 {
					y -= 1
				} else {
					yr -= 1
				}
			}
			if ch == '>' {
				if i%2 == 0 {
					x += 1
				} else {
					xr += 1
				}
			}
			if ch == '<' {
				if i%2 == 0 {
					x -= 1
				} else {
					xr -= 1
				}
			}

			if i%2 == 0 {
				res[fmt.Sprintf("%d,%d", x, y)] = true
			} else {
				res[fmt.Sprintf("%d,%d", xr, yr)] = true
			}
		}
		result += len(res)
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

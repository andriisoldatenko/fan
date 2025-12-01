package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// to handle negative mod
func mod(a, b int) int {
	return (a%b + b) % b
}

func div(a, b int) int {
	if a < 0 {
		return -((-a + b - 1) / b)
	}
	return a / b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 50
	rr := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text() // GET the line string
		t = strings.Replace(t, "L", "-", -1)
		t = strings.Replace(t, "R", "", -1)
		number, _ := strconv.Atoi(t)
		start_zero := result == 0

		d := div(result+number, 100)
		m := mod(result+number, 100)

		result = m

		rr = rr + abs(d)

		if number < 0 {
			rr = rr + boolToInt(result == 0) - boolToInt(start_zero)
		}

	}
	fmt.Println(rr)
}

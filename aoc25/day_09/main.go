package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func sum_pos(arr []int) int {
	result := 0
	for _, a := range arr {
		if a > 0 {
			result += a
		}
	}
	return result
}

func print_(obj any) { fmt.Println(obj) }

func printM(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

func printArr(a [10][10]string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Coord struct {
	x int
	y int
}

func (a Coord) area(b Coord) int {
	return (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
}

func part1() {
	path := "/Users/andrii.soldatenko/work/fan/aoc25/day_09/input.txt"
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rawData := strings.Split(string(file), "\n")

	coords := []Coord{}
	// we can skip every second line
	for i := 0; i < len(rawData); i++ {
		dd := strings.Split(rawData[i], ",")
		coords = append(coords, Coord{x: atoi(dd[0]), y: atoi(dd[1])})
	}

	maxArea := -1
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			maxArea = max(coords[i].area(coords[j]), maxArea)
		}
	}
	print_(maxArea)

}

func main() {
	part1()
	//part2()
}

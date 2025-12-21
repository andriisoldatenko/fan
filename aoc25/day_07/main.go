package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// BFS graph

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

type Split struct {
	x1 int
	x2 int
}

func part1() {

	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc25/day_07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	puzzle := [][]string{}
	splits := []Split{}

	i := 0
	for sc.Scan() {
		t := sc.Text()
		ll := strings.Split(t, "")
		for j, ch := range ll {
			if ch == "^" {
				splits = append(splits, Split{x1: i, x2: j})
			}
		}
		puzzle = append(puzzle, ll)
		i++
	}

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == "S" {
				puzzle[i+1][j] = "|"
				break
			}
		}
	}

	for _, split := range splits {
		leftStop := false
		rightStop := false
		for i := split.x1; i < len(puzzle); i++ {
			if leftStop && rightStop {
				break
			}
			if puzzle[i][split.x2-1] == "." && !leftStop {
				puzzle[i][split.x2-1] = "|"
			}
			if puzzle[i][split.x2+1] == "." && !rightStop {
				puzzle[i][split.x2+1] = "|"
			}
			if puzzle[i][split.x2-1] == "^" {
				leftStop = true
			}
			if puzzle[i][split.x2+1] == "^" {
				rightStop = true
			}
		}
		//printM(puzzle)
	}

	//printM(puzzle)
	//print_(splits)
	result := 0
	for _, split := range splits {
		if puzzle[split.x1-1][split.x2] == "|" {
			result += 1
		}
	}
	print_(result)
}

func atoi(num string) int {
	n, _ := strconv.Atoi(num)
	return n
}

func part2() {
	path := "/Users/andrii.soldatenko/work/fan/aoc25/day_07/input.txt"
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rawData := strings.Split(string(file), "\n")

	data := []string{}
	// we can skip every second line
	for i := 0; i < len(rawData); i += 2 {
		data = append(data, rawData[i])
	}

	processedRows := [][]int{}
	firstLine := make([]int, len(data[0]))

	for i, ch := range data[0] {
		if ch == 'S' {
			firstLine[i] = 1
		}
	}

	processedRows = append(processedRows, firstLine)

	splits := 0
	for _, row := range data[1:] {

		current := make([]int, len(row))
		for ri, ch := range row {
			if ch == '^' {
				current[ri] = -1
			}
		}

		for i := 1; i < len(row); i++ {
			char := row[i]
			above := processedRows[len(processedRows)-1][i]
			if above > 0 {
				if char == '^' {
					splits += 1
					current[i-1] += above
					current[i+1] += above
				} else {
					current[i] += above
				}
			}
		}

		processedRows = append(processedRows, current)
	}
	printM(processedRows)
	print_(splits)
	print_(sum_pos(processedRows[len(processedRows)-1]))
}

func main() {
	//part1()
	part2()
}

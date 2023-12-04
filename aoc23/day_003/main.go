package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	start     int
	end       int
	row       int
	value     int
	asterisks []Asterisk
}

type Asterisk struct {
	i int
	j int
}

const MAX int = 140

func part1() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_003/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	bb := make([][]rune, MAX)
	for i := 0; i < MAX; i++ {
		bb[i] = make([]rune, MAX)
	}

	row := 0
	allNumbers := []Number{}
	for sc.Scan() {
		t := sc.Text()
		for index, ch := range t {
			bb[row][index] = ch
		}
		allNumbers = append(allNumbers, findNumbers(t, row)...)
		row += 1
	}

	bb = append([][]rune{[]rune(strings.Repeat(".", MAX))}, bb...)
	bb = append(bb, []rune(strings.Repeat(".", MAX)))

	for index := range bb {
		bb[index] = insert(bb[index], 0, '.')
		bb[index] = append(bb[index], '.')
	}
	result := 0

	for _, number := range allNumbers {
		if includeNumber(number.start+1, number.end+1, number.row+1, bb) {
			result += number.value
		}
	}
	fmt.Println(result)
}

func main() {
	//part1()
	part2()
}

func part2() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_003/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	bb := make([][]rune, MAX)
	for i := 0; i < MAX; i++ {
		bb[i] = make([]rune, MAX)
	}

	row := 0
	allNumbers := []Number{}
	for sc.Scan() {
		t := sc.Text()
		for index, ch := range t {
			bb[row][index] = ch
		}

		allNumbers = append(allNumbers, findNumbers(t, row)...)
		row += 1
	}

	bb = append([][]rune{[]rune(strings.Repeat(".", MAX))}, bb...)
	bb = append(bb, []rune(strings.Repeat(".", MAX)))

	for index := range bb {
		bb[index] = insert(bb[index], 0, '.')
		bb[index] = append(bb[index], '.')
	}

	updated := []Number{}
	for _, number := range allNumbers {
		num := addStar(bb, number)
		//fmt.Println(num)
		updated = append(updated, num)
	}

	result := map[string][]int{}
	for _, n1 := range updated {
		for _, a := range n1.asterisks {
			result[fmt.Sprintf("%d%d", a.i, a.j)] = append(result[fmt.Sprintf("%d%d", a.i, a.j)], n1.value)
		}
	}

	sum := 0
	for k := range result {
		if len(result[k]) > 1 {
			sum += Product(result[k])
		}
	}
	fmt.Println(sum)

}

func Product(arr []int) int {
	r := 1
	for _, a := range arr {
		r *= a
	}
	return r
}

func findNumbers(a string, row int) []Number {
	start := 0
	end := 0
	num := []string{}
	numbers := []Number{}
	for index, item := range a {
		if unicode.IsDigit(item) {
			if start == -1 {
				start = index
			}
			num = append(num, string(item))
			end = index
		} else {
			if start != -1 {
				value, err := strconv.Atoi(strings.Join(num, ""))
				if err == nil {
					numbers = append(numbers, Number{
						start: start,
						end:   end,
						row:   row,
						value: value,
					})
				}
				//fmt.Println(strings.Join(num, ""))
				//fmt.Println(start, end)
			}
			num = []string{}
			start = -1
			end = -1
		}

		if index == len(a)-1 && unicode.IsDigit(item) {
			value, err := strconv.Atoi(strings.Join(num, ""))
			if err == nil {
				numbers = append(numbers, Number{
					start: start,
					end:   end,
					row:   row,
					value: value,
				})
			}
			//fmt.Println(strings.Join(num, ""))
			//fmt.Println(start, end)
		}
	}
	return numbers
}

func addStar(board [][]rune, number Number) Number {
	// row = i, start j
	// row = i, end j
	asterisks := []Asterisk{}

	start := number.start + 1
	end := number.end + 1
	row := number.row + 1

	left := board[row][start-1]
	if left == '*' {
		asterisks = append(asterisks, Asterisk{
			i: row, j: start - 1,
		})
	}
	right := board[row][end+1]
	if right == '*' {
		asterisks = append(asterisks, Asterisk{
			i: row, j: end + 1,
		})
	}

	for i := start - 1; i <= end+1; i++ {
		tm := board[row-1][i]
		if tm == '*' {
			asterisks = append(asterisks, Asterisk{
				i: row - 1, j: i,
			})
		}
	}

	for i := start - 1; i <= end+1; i++ {
		bm := board[row+1][i]
		if bm == '*' {
			asterisks = append(asterisks, Asterisk{
				i: row + 1, j: i,
			})
		}
	}

	number.asterisks = asterisks
	return number
}

func includeNumber(start, end, row int, board [][]rune) bool {
	// row = i, start j
	// row = i, end j
	result := []rune{}

	result = append(result, board[row][start-1])
	result = append(result, board[row][end+1])

	for i := start - 1; i <= end+1; i++ {
		tm := board[row-1][i]
		result = append(result, tm)
	}

	for i := start - 1; i <= end+1; i++ {
		bm := board[row+1][i]
		result = append(result, bm)
	}

	if allSameStrings(result) && result[0] == '.' {
		return false
	}
	fmt.Println(result)
	return true
}

func allSameStrings(a []rune) bool {
	for _, v := range a {
		if v != a[0] {
			return false
		}
	}
	return true
}

func insert(a []rune, index int, value rune) []rune {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func PrintBoard(bb [][]rune) {
	for _, row := range bb {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

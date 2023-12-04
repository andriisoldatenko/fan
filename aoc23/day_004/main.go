package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func part1() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_004/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")
	result := 0
	for sc.Scan() {
		t := sc.Text()
		row := strings.Split(strings.Split(t, ":")[1], "|")
		lucky := re.FindAllString(row[0], -1)
		your := re.FindAllString(row[1], -1)

		sum := 0
		for _, l := range lucky {
			for _, y := range your {
				if l == y {
					if sum == 0 {
						sum = 1
					} else {
						sum *= 2
					}
				}
			}
		}
		result += sum
	}
	fmt.Println(result)
}

func main() {
	//part1()
	part2()
}

var result []int

func part2() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_004/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")
	rowN := 1
	cards := map[int][]int{}
	for sc.Scan() {
		t := sc.Text()
		row := strings.Split(strings.Split(t, ":")[1], "|")
		lucky := re.FindAllString(row[0], -1)
		your := re.FindAllString(row[1], -1)

		i := 0
		for _, l := range lucky {
			for _, y := range your {
				if l == y {
					i++
				}
			}
		}
		for j := rowN + 1; j < i+rowN+1; j++ {
			cards[rowN] = append(cards[rowN], j)
		}
		rowN++
	}
	for i := 1; i <= 208; i++ {
		walk(cards, i)
	}
	fmt.Println(len(result))

}

func walk(cards map[int][]int, val int) {
	if v, ok := cards[val]; ok {
		result = append(result, val)
		for _, k := range v {
			walk(cards, k)
		}
	} else {
		result = append(result, val)
		return
	}
}

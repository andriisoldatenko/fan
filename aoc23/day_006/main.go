package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func part1() {
	file, err := os.ReadFile("/Users/andrii.soldatenko/work/fan/aoc23/day_006/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)

	all := strings.Split(content, "\n")
	times := strToints(all[0])
	distances := strToints(all[1])
	fmt.Println(times)
	fmt.Println(distances)

	//v := 1
	r := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		// speed = distance / time

		result := []int{}
		for s := 1; s <= time; s++ {
			d := (time - s) * s
			if d > distance {
				result = append(result, s)
			}
		}
		//fmt.Println(result)
		r *= len(result)
	}
	fmt.Println(r)
}

func part2() {
	file, err := os.ReadFile("/Users/andrii.soldatenko/work/fan/aoc23/day_006/input_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)

	all := strings.Split(content, "\n")
	times := strToints(all[0])
	distances := strToints(all[1])
	fmt.Println(times)
	fmt.Println(distances)

	////v := 1
	//r := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		// speed = distance / time

		result := []int{}
		for s := 1; s <= time; s++ {
			d := (time - s) * s
			if d > distance {
				result = append(result, s)
			}
		}
		fmt.Println(len(result))
		//r *= len(result)
	}
	//fmt.Println(r)

}

func main() {
	//part1()
	part2()
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

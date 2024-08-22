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
		t := sc.Text() // GET the line string
		res := strToints(t)
		l := res[0]
		w := res[1]
		h := res[2]
		result += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
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
		res := strToints(t)
		l := res[0]
		w := res[1]
		h := res[2]
		result += l*w*h + min(2*l+2*w, 2*w+2*h, 2*l+2*h)
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

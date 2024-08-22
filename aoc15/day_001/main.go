package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Println(t)
		for _, ch := range t {
			if ch == '(' {
				result++
			} else {
				result--
			}

		}
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
		fmt.Println(t)
		for index, ch := range t {
			if ch == '(' {
				result++
			} else {
				result--
			}
			if result == -1 {
				fmt.Println(index + 1)
				break
			}
		}
	}
	fmt.Println(result)
}

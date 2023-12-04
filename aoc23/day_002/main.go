package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//var possibleGames = map[string]int{
//	"red":   12,
//	"green": 13,
//	"blue":  14,
//}
//
//func main() {
//	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_002/input.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	sc := bufio.NewScanner(file)
//	result := 0
//	for sc.Scan() {
//		t := sc.Text()
//		res := strings.Split(t, ":")
//		rr := strings.Split(res[0], " ")
//		game, _ := strconv.Atoi(rr[1])
//		games := strings.Split(res[1], ";")
//		isPossible := true
//		for _, g := range games {
//			g = strings.TrimLeft(g, " ")
//			sets := strings.Split(g, ", ")
//			for _, s := range sets {
//				p := strings.Split(s, " ")
//				n, _ := strconv.Atoi(p[0])
//				color := p[1]
//				val, _ := possibleGames[color]
//				if n > val {
//					isPossible = false
//					break
//				}
//			}
//			if !isPossible {
//				break
//			}
//		}
//		if isPossible {
//			result += game
//		}
//
//	}
//	fmt.Println(result)
//}

func main() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_002/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sum := 0
	for sc.Scan() {
		t := sc.Text()
		result := map[string]int{"red": 0, "blue": 0, "green": 0}
		res := strings.Split(t, ":")
		//rr := strings.Split(res[0], " ")
		//game, _ := strconv.Atoi(rr[1])
		games := strings.Split(res[1], ";")
		for _, g := range games {
			g = strings.TrimLeft(g, " ")
			sets := strings.Split(g, ", ")
			for _, s := range sets {
				p := strings.Split(s, " ")
				n, _ := strconv.Atoi(p[0])
				color := p[1]
				val, _ := result[color]
				if n > val {
					result[color] = n
				}
			}
		}
		//
		product := 1
		for k := range result {
			product = product * result[k]
		}
		sum += product
		fmt.Println(product)
	}
	fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"strings"
)

//type Edge struct {
//	source, target int
//}
//
//func Bellman_ford(edges []Edge, num_vertices, source int, target int) ([]int, error) {
//
//	distances := make([]int, num_vertices)
//	for i := range distances {
//		distances[i] = math.MaxInt32
//	}
//	distances[source] = 0
//
//	for i := 0; i < num_vertices-1; i++ {
//		for _, edge := range edges {
//			if distances[edge.source]+edge.weight < distances[edge.target] {
//				distances[edge.target] = distances[edge.source] + edge.weight
//			}
//		}
//	}
//
//	for _, edge := range edges {
//		if distances[edge.source]+edge.weight < distances[edge.target] {
//			return nil, fmt.Errorf("graph contains negative-weight cycle")
//		}
//	}
//	return distances, nil
//}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func part1() {
	file, err := os.ReadFile("/Users/andrii.soldatenko/work/fan/aoc23/day_008/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)

	all := strings.Split(content, "\n")
	instructions := all[0]
	fmt.Println(instructions)

	m := map[string][]string{}

	start := []string{}

	for _, line := range all[2:] {
		ll := line[:3]
		m[ll] = []string{line[7:10], line[12:15]}
		if ll[len(ll)-1] == 'A' {
			start = append(start, ll)
		}
		//fmt.Println(line)
	}
	//fmt.Println(start)

	steps := 1
	current := start[0]
	// 20803
	// 17873
	// 23147
	// 15529
	// 17287
	// 19631

out:
	for {
		for _, ch := range instructions {
			if ch == 'R' {
				val := m[current]
				current = val[1]
			}

			if ch == 'L' {
				val := m[current]
				current = val[0]
			}

			if current[len(current)-1] == 'Z' {
				break out
			}
			steps += 1
		}

	}
	//
	fmt.Println(steps)

}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func part2() {
	//
	// 20803
	// 17873
	// 23147
	// 15529
	// 17287
	// 19631
	fmt.Println(LCM(20803, 17873, 23147, 15529, 17287, 17287, 19631))
}

func allEndsWithZ(in []string) bool {
	for _, i := range in {
		if i[len(i)-1] != 'Z' {
			return false
		}
	}
	return true
}

func main() {

	part1()
	part2()
}

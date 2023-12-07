package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	file, err := os.ReadFile("/Users/andrii.soldatenko/work/fan/aoc23/day_005/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)

	all := strings.Split(content, "\n\n")

	seedsS := strings.Split(all[0][7:], " ")
	seeds := make([]int, len(seedsS))
	for i, s := range seedsS {
		seeds[i], _ = strconv.Atoi(s)
	}

	seedToSoil := parse(all[1])
	fmt.Println(seedToSoil)

	soilToFertilizer := parse(all[2])
	fmt.Println(soilToFertilizer)

	fertilizerToWater := parse(all[3])
	fmt.Println(fertilizerToWater)

	waterToLight := parse(all[4])
	fmt.Println(waterToLight)

	lightToTemperature := parse(all[5])
	fmt.Println(lightToTemperature)

	temperatureToHumidity := parse(all[6])
	fmt.Println(temperatureToHumidity)

	humidityToLocation := parse(all[7])
	fmt.Println(humidityToLocation)

	chunks := chunkBy(seeds, 2)

	result := []int{}
	for _, chunk := range chunks {
		//newSeeds = append(newSeeds, chunk[0])
		//newSeeds = append(newSeeds, chunk[0]+chunk[1]-1)
		start := chunk[0]
		end := chunk[0] + chunk[1] - 1

		i := 0
		localRes := []int{}
		for s := start; s <= end; s++ {
			num := process(s, seedToSoil)
			num = process(num, soilToFertilizer)
			num = process(num, fertilizerToWater)
			num = process(num, waterToLight)
			num = process(num, lightToTemperature)
			num = process(num, temperatureToHumidity)
			num = process(num, humidityToLocation)
			//fmt.Printf("seed: %d -> %d\n", s, num)
			localRes = append(localRes, num)

			if len(localRes) > 2 {
				if localRes[i-1] < num && localRes[i-2] > localRes[i-1] {
					result = append(result, localRes[i-1])
				}
			}
			i++
		}
		fmt.Println(result)
	}

	m := 0
	for i, v := range result {
		if i == 0 || v < m {
			m = v
		}
	}
	fmt.Println(m)

	//m := 0
	//i := 0
	//for _, s := range newSeeds {
	//	num := process(s, seedToSoil)
	//	num = process(num, soilToFertilizer)
	//	num = process(num, fertilizerToWater)
	//	num = process(num, waterToLight)
	//	num = process(num, lightToTemperature)
	//	num = process(num, temperatureToHumidity)
	//	num = process(num, humidityToLocation)
	//
	//	if i == 0 || num < m {
	//		m = num
	//	}
	//	fmt.Printf("seed: %d -> %d\n", s, num)
	//	i++
	//}

	//fmt.Println(m)
}

func part2() {

}

func main() {
	part1()
	part2()
}

func convert(src []string) []int {
	result := make([]int, len(src))
	for i, s := range src {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

func parse(src string) [][]int {
	w := strings.Split(src, "\n")[1:]
	f := make([][]int, len(w))
	for i := range w {
		f[i] = append(f[i], convert(strings.Split(w[i], " "))...)
	}
	return f
}

func process(s int, m [][]int) int {
	found := false
	result := 0
	for _, row := range m {
		dest := row[0]
		src := row[1]
		length := row[2]
		if (src+length-1 >= s) && (s >= src) {
			found = true
			result = s - src + dest
		}
		if found {
			break
		}
	}
	if !found {
		result = s
	}
	return result
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

//
//
//seed-to-soil map:
//50 98 99 -> seed - 48
//52 50 97 -> seed + 2
// -> seed

//0 15 41 -> seed + 0 - 15
//37 52 53 -> seed + 37 - 52
//39 0 15

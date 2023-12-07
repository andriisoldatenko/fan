package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"slices"
	"sort"
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

type Type int

const (
	FiveKind Type = 7 - iota
	FourKind
	FullHouse
	ThreeKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	cards string
	type_ Type
	num   int
}

func (h *Hand) Strength() {
	counter := map[rune]int{}
	for _, c := range h.cards {
		counter[c]++
	}

	if len(counter) == 1 {
		h.type_ = FiveKind
		return
	}

	vals := Values(counter)

	if reflect.DeepEqual(vals, []int{1, 4}) {
		if _, ok := counter[74]; ok {
			h.type_ = FiveKind
			return
		}
		h.type_ = FourKind
		return
	}

	if reflect.DeepEqual(Values(counter), []int{2, 3}) {
		if _, ok := counter[74]; ok {
			h.type_ = FiveKind
			return
		}
		h.type_ = FullHouse
		return
	}

	if reflect.DeepEqual(Values(counter), []int{1, 1, 3}) {
		if _, ok := counter[74]; ok {
			h.type_ = FourKind
			return
		}
		h.type_ = ThreeKind
		return
	}

	if reflect.DeepEqual(Values(counter), []int{1, 2, 2}) {
		if val, ok := counter[74]; ok {
			if val == 2 {
				h.type_ = FourKind
				return
			} else if val == 1 {
				h.type_ = FullHouse
				return
			}
		}
		h.type_ = TwoPair
		return
	}

	if reflect.DeepEqual(Values(counter), []int{1, 1, 1, 2}) {
		if _, ok := counter[74]; ok {
			h.type_ = ThreeKind
			return
		}
		h.type_ = OnePair
		return
	}

	if _, ok := counter[74]; ok {
		h.type_ = OnePair
		return
	}

	h.type_ = HighCard

}

type kv struct {
	Key   rune
	Value int
}

func Values(a map[rune]int) []int {
	result := []int{}
	for k := range a {
		result = append(result, a[k])
	}
	slices.Sort(result)
	return result
}

// A - 65 -> 65
// K - 75 -> 64
// Q - 81 -> 63
// J - 74 -> 62
// T - 84 -> 61
// 9 - 57 -> 57
// 2 - 50 -> 50
// J - 74 -> 49
var mapping = map[uint8]uint8{
	75: 64,
	81: 63,
	//74: 62, // -- part 1
	84: 61,
	74: 49, // -- part 2
}

func part1() {
	file, err := os.Open("/Users/andrii.soldatenko/work/fan/aoc23/day_007/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	hands := []Hand{}
	for sc.Scan() {
		t := sc.Text()
		split := strings.Fields(t)
		h := split[0]

		num, _ := strconv.Atoi(split[1])
		hand := Hand{cards: h, num: num}
		hand.Strength()
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].type_ != hands[j].type_ {
			return hands[i].type_ > hands[j].type_
		}

		for index := 0; index < 5; index++ {
			if hands[i].cards[index] != hands[j].cards[index] {
				left := hands[i].cards[index]
				right := hands[j].cards[index]

				if v, ok := mapping[left]; ok {
					left = v
				}

				if v, ok := mapping[right]; ok {
					right = v
				}

				return left > right
			}
		}

		return false
	})

	sum := 0
	for i, h := range hands {
		sum = sum + (h.num * (len(hands) - i))
	}
	fmt.Println(sum)
	fmt.Println(sum)

}

func part2() {

}

func main() {

	part1()
	//part2()
}

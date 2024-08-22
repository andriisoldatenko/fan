package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
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
		t := sc.Text()
		for i := 1; i < 1_000_0000; i++ {
			r := GetMD5Hash(fmt.Sprintf("%s%d", t, i))
			if strings.HasPrefix(r, "000000") {
				fmt.Println(i)
				break
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
		t := sc.Text()
		fmt.Println(t)
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

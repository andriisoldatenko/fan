package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	line := readLine(reader)

	nums := strings.Split(line, " ")
	ints := make([]int, len(nums))
	max := 0
	for i, s := range nums {
		ints[i], _ = strconv.Atoi(s)
		if ints[i] > max {
			max = ints[i]
		}
	}
	var results []string
	for i := 0; i < len(ints); i++ {
		if max-ints[i] != 0 {
			results = append(results, strconv.Itoa(max-ints[i]))
		}
	}
	fmt.Println(strings.Join(results, " "))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

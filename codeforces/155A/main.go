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
	readLine(reader)
	line := readLine(reader)

	nums := strings.Split(line, " ")
	ints := make([]int, len(nums))
	for i, s := range nums {
		ints[i], _ = strconv.Atoi(s)
	}
	n := len(ints)
	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			fmt.Print(ints[j])
		}
		fmt.Println()
	}
	fmt.Println(ints)
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

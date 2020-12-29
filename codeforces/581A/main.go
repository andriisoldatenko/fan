package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	line := readLine(reader)
	var a, b int64 = 0, 0
	fmt.Sscanf(line, "%d %d", &a, &b)
	fmt.Println(min(a, b), abs(a-b)/2)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

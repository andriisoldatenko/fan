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
	var n, m int = 0, 0
	fmt.Sscanf(line, "%d %d", &n, &m)
	swap := true
	for i := 0; i < n; i++ {
		line := ""
		for j := 0; j < m; j++ {
			if i%2 != 0 {
				line += "."
			} else {
				line += "#"
			}
		}
		if i%2 != 0 {
			if swap {
				line = line[:m-1] + "#"
			} else {
				line = "#" + line[1:]
			}
			swap = !swap
		}
		fmt.Println(line)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	line := readLine(reader)
	t := 0
	fmt.Sscanf(line, "%d", &t)
	for i := 0; i < t; i++ {
		var d, res int64
		fmt.Sscanf(readLine(reader), "%d", &d)
		if d % 2 == 0 {
			res = (d / 2) - 1
		} else {
			res = d / 2
		}
		fmt.Println(res)
	}
}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
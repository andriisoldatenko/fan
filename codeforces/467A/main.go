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
	n := 0
	fmt.Sscanf(readLine(reader), "%d", &n)
	res := 0
	for i := 0; i < n; i++ {
		var p, q int
		fmt.Sscanf(readLine(reader), "%d %d", &p, &q)
		if q - p >= 2 {
			res++
		}
	}
	fmt.Println(res)
}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
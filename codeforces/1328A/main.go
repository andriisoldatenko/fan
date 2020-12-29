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

	for i := 0; i < n; i++ {
		line := readLine(reader)
		a,b := 0, 0
		fmt.Sscanf(line, "%d %d", &a, &b)
		if a % b == 0 {
			fmt.Println(0)
			continue
		}
		d := ((a / b) + 1)*b
		fmt.Println(d - a)
	}
}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
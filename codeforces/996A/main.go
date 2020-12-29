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
	var n int64 = 0
	fmt.Sscanf(line, "%d", &n)
	var count int64 = 0
	for {
		if n < 100 {
			break
		}
		n = n - 100
		count++
	}
	for {
		if n < 20 {
			break
		}
		n = n - 20
		count++
	}
	for {
		if n < 10 {
			break
		}
		n = n - 10
		count++
	}

	for {
		if n < 5 {
			break
		}
		n = n - 5
		count++
	}
	count = count + n
	fmt.Println(count)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

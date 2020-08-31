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
	var k, n, w int
	fmt.Sscanf(readLine(reader), "%d %d %d", &k, &n, &w)

	total := 0
	for i := 1; i <= w; i++ {
		total += k * i
	}
	if n >= total {
		fmt.Println(0)
	} else {
		fmt.Println(total-n)
	}
}



func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
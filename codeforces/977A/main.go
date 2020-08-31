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
	var n, k int
	fmt.Sscanf(readLine(reader), "%d %d", &n, &k)

	for i := 0; i < k; i++ {
		if n % 10 == 0 {
			n /= 10
		} else {
			n--
		}
	}
	fmt.Println(n)
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
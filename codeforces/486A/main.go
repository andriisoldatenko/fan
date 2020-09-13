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
	reader := bufio.NewReaderSize(os.Stdin, 2048 * 2048)
	n, _ := strconv.ParseInt(readLine(reader), 10, 64)
	if n % 2 == 0 {
		fmt.Print(n / 2)
	} else {
		fmt.Print(-(n+1)/2)
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
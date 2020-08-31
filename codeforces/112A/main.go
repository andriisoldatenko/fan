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
	lineA := readLine(reader)
	lineB := readLine(reader)
	res := 0
	for i := 0; i < len(lineA); i++ {
		c1 := strings.ToLower(string(lineA[i]))
		c2 := strings.ToLower(string(lineB[i]))
		if c1[0] < c2[0] {
			res = -1
			break
		}
		if c1[0] > c2[0] {
			res = 1
			break
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
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
	var n int
	fmt.Sscanf(readLine(reader), "%d", &n)

	out := strings.Split(readLine(reader), " ")

	for _, c := range out {
		if string(c) == "1" {
			fmt.Println("HARD")
			return
		}
	}
	fmt.Println("EASY")
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
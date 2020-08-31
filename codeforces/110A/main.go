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
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	var n int

	luckyNums := 0
	out := readLine(reader)

	for _, c := range out {
		if string(c) == "4" || string(c) == "7" {
			luckyNums++
		}
	}

	s := strconv.Itoa(luckyNums)

	fmt.Println(n)
}


func isLucky(s string) {

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
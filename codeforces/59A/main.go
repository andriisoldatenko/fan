package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	out := readLine(reader)
	lowerCount := 0
	upperCount := 0
	for _, c := range out {
		if unicode.IsLower(c) {
			lowerCount++
		} else {
			upperCount++
		}
	}

	if lowerCount >= upperCount {
		fmt.Println(strings.ToLower(out))
	} else {
		fmt.Println(strings.ToUpper(out))
	}
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
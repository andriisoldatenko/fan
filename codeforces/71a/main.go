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
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)

	checkError(err)
	x := int(nTemp)
	for i := 0; i < x; i++ {
		word := readLine(reader)
		if len(word) <= 10 {
			fmt.Println(word)
		} else {
			fmt.Printf("%s%d%s\n", string(word[0]), len(word) - 2, string(word[len(word) - 1]))
		}
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
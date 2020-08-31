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
	res := 0
	for i := 0; i < x; i++ {
		out := readLine(reader)
		if out == "++X" || out == "X++" {
			res++
		} else {
			res--
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
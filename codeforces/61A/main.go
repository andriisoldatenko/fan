package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2048 * 2048)
	n1 := readLine(reader)
	n2 := readLine(reader)
	res := ""
	for i:=0; i<len(n1); i++ {
		if n1[i] != n2[i]{
			res += "1"
		} else {
			res += "0"
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
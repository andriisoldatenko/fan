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
	var s []string
	for i:=1; i<=n; i++ {
		if i % 2 != 0 {
			s = append(s, "I hate")
		} else {
			s = append(s, "I love")
		}
	}
	res := strings.Join(s, " that ")
	fmt.Printf("%s it\n", res)
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
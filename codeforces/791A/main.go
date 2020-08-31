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
	var a, b int
	fmt.Sscanf(readLine(reader), "%d %d", &a, &b)

	year := 0
	limak := a
	bob := b
	for {
		if limak > bob {
			fmt.Println(year)
			break
		}
		year++
		limak = 3 * limak
		bob = 2 * bob
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
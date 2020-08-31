package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	ix, jx := 0, 0
	for i :=0; i < 5; i++ {
		out := strings.Split(readLine(reader), " ")
		for j, c := range out {
			if c == "1" {
				ix = i
				jx = j
				fmt.Println(math.Abs(float64(2-ix)) + math.Abs(float64(2-jx)))
				break
			}
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
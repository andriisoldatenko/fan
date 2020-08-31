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
	n := 0
	fmt.Sscanf("%d", readLine(reader), &n)
	_ = n
	row := readLine(reader)
	a := 0
	d := 0
	for _, c := range row {
		if string(c) == "A" {
			a++
		} else {
			d++
		}
	}
	if a == d {
		fmt.Println("Friendship")
	}
	if a > d {
		fmt.Println("Anton")
	} else if a < d {
		fmt.Println("Danik")
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
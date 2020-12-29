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
	line := readLine(reader)
	out := strings.Split(line[1:len(line)-1], ", ")
	set := make(map[string]bool)
	for _, c := range out {
		if c != "" {
			set[c] = true
		}
	}
	fmt.Println(len(set))
}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
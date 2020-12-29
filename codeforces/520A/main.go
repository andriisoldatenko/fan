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
	fmt.Sscanf(readLine(reader), "%d", &n)
	set := make(map[string]bool)
	if n <= 25 {
		fmt.Println("NO")
	} else {
		line := readLine(reader)
		for _,c := range line {
			set[strings.ToLower(string(c))] = true
		}
		if len(set) >= 26 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
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
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
	yearStr := readLine(reader)
	year := 0
	fmt.Sscanf(yearStr, "%d", &year)
	year = year + 1
	for i := year; i <= 90000; i++ {
		set := make(map[string]bool)
		for _,c := range strconv.Itoa(i) {
			set[string(c)] = true
		}
		if len(set) == 4 {
			fmt.Println(i)
			break
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
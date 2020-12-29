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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	t := 0
	fmt.Sscanf(readLine(reader), "%d", &t)
	for i := 0; i < t; i++ {
		num := 0
		fmt.Sscanf(readLine(reader), "%d", &num)
		r := solve(num)
		fmt.Println(len(r))
		fmt.Println(strings.Join(r, " "))
	}

}

func solve(value int) []string {
	var results []string
	digit := 0
	i := 1
	for value > 0 {
		digit = value % 10
		if digit*i > 0 {
			results = append(results, strconv.Itoa(digit*i))
		}
		value /= 10
		i *= 10
	}
	return results
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

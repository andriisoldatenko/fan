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

	set := map[string]int{
		"Tetrahedron": 4,
		"Cube": 6,
		"Octahedron": 8,
		"Dodecahedron": 12,
		"Icosahedron": 20,
	}
	sum := 0
	for i := 0; i < n; i++ {
		line := readLine(reader)
		sum+=set[line]
	}
	fmt.Println(sum)
}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)


func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	names := SortString(readLine(reader) + readLine(reader))
	concat := SortString(readLine(reader))


	if names == concat {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}


func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
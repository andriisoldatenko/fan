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
	var n, k int
	fmt.Sscanf(readLine(reader), "%d %d", &n, &k)
	out := strings.Split(readLine(reader), " ")
	ints := make([]int, len(out))

	for i, s := range out {
		ints[i], _ = strconv.Atoi(s)
	}
	res := 0
	for _, i := range ints {
		if i >= ints[k-1] && i > 0 {
			res += 1
		}
	}
	fmt.Println(res)
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
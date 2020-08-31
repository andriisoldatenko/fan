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
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)

	checkError(err)
	x := int(nTemp)
	total := 0
	for i := 0; i < x; i++ {
		out := strings.Split(readLine(reader), " ")
		ints := make([]int, len(out))

		for i, s := range out {
			ints[i], _ = strconv.Atoi(s)
		}
		sum := 0
		for _, i := range ints {
			sum += i
		}
		if sum >= 2 {
			total += 1
		}
	}
	fmt.Println(total)
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
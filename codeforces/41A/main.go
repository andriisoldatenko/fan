package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


func isPalindrome(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var j = len(s1) - 1
	for i := 0; i <= j; i++ {
		if s1[j-i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	word := readLine(reader)
	word2 := readLine(reader)
	if isPalindrome(word, word2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}


func isLucky(s string) {

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
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func subStringsFromString(s string) []string {
	var result []string
	for i := 0; i < len(s); i++ {
		for j := i+1; j <= len(s); j++ {
			result = append(result, s[i:j])
		}
	}
	return result
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Complete the sherlockAndAnagrams function below.
//func sherlockAndAnagramsv1(s string) int32 {
//	r1 := 0
//
//	dupFreq := make(map[string]int)
//	letters := strings.Split(s, "")
//	for _, char := range letters {
//		dupFreq[char] += 1
//	}
//
//	if len(s) == len(dupFreq) {
//		return 0
//	}
//
//	subs := subStringsFromString(s)
//	fmt.Println(subs)
//	for i := 0; i < len(subs); i++ {
//		for j := i+1; j < len(subs); j++ {
//			if sortString(subs[i]) == sortString(subs[j]) {
//				r1 += 1
//			}
//		}
//	}
//	return int32(r1)
//}


// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	n := len(s)
	arr := strings.Split(s, "")
	mp := make(map[string]int)

	for i := 0; i < n; i++ {
		sb := ""
		for j := i; j < n; j++ {
			s1 := fmt.Sprint(sb, arr[j])
			sb = sortString(s1)
			fmt.Println(sb)
			mp[sb] += 1
		}
	}
	r1 := 0
	for _, v := range mp {
		//fmt.Println(k, v)
		r1 = r1 + (v * ( v - 1 ) / 2)
	}
	return int32(r1)
}


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//defer stdout.Close()
	//writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Println(result)
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

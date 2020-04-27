package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func prettyPrint(arr [][]int32) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	maxSum := int32(math.MinInt32)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var newArr [][]int32
			for k := 0; k < 3; k++ {
				var arrRow []int32
				for l := 0; l < 3; l++ {
					arrRow = append(arrRow, arr[k+i][l+j])
				}
				newArr = append(newArr, arrRow)
			}
			//prettyPrint(newArr)
			maxSum = max(subSum(newArr), maxSum)
			//fmt.Println(maxSum)
		}
	}
	return maxSum
}


func subSum(arr [][]int32) int32 {
	var result int32 = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i == 1 && j == 0) || (i == 1 && j == 2) {
				continue
			}
			result += arr[i][j]
		}
	}

	return result
}


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//defer stdout.Close()
	//writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)
	fmt.Print(result)
	//fmt.Fprintf(writer, "%d\n", result)

	//writer.Flush()
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

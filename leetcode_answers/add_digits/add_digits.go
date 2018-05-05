package main

import (
	"fmt"
)
func splitNumberToSlice(num int) []int {
	var arr []int
	for {
		arr = append(arr, num % 10)
		num /= 10
		if num > 0 {
			continue
		} else { break }
	}
	return arr
}


func addDigitsSimple(num int) int {
	nums := splitNumberToSlice(num)
	result := 0
	for _, n := range nums {
		result += n
	}
	if result < 10 {
		return result
	} else {
		return addDigitsSimple(result)
	}
	return result
}

func addDigits(num int) int {
	return 1 + ((num - 1) % 9)
}

func main() {
	fmt.Println(addDigits(38))
}
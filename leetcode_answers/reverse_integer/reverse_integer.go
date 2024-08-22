package main

import (
	"fmt"
	"math"
)

// https://golang.org/ref/spec#Numeric_types
const MaxInt = 2147483647
const MinInt = -MaxInt - 1

func splitNumberToSlice(num int) []int {
	var arr []int
	for {
		arr = append(arr, num%10)
		num /= 10
		if num > 0 {
			continue
		} else {
			break
		}
	}
	return arr
}

func reverseInt(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	nums := reverseInt(splitNumberToSlice(x))
	result := 0
	for i, n := range nums {
		result = result + n*int(math.Pow10(i))
	}
	if negative {
		result = -result
	}

	if result > MaxInt || result < MinInt {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(-2147483648))
}

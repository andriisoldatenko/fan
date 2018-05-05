package main

import "fmt"

func All(vs []bool) bool{
	for _, v := range vs {
		if !v {
			return false
		}
	}
	return true
}

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

func selfDividingNumbers(left int, right int) []int {
	var results []int
	for i := left; i <= right; i++ {
		var selfDiv []bool
		for _, c := range splitNumberToSlice(i) {
			if c == 0 {
				selfDiv = append(selfDiv, false)
			} else {
				selfDiv = append(selfDiv, i % c == 0)
			}
		}
		if All(selfDiv) {
			results = append(results, i)
		}
	}
	return results
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
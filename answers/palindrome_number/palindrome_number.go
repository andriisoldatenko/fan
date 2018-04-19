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


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	nums := splitNumberToSlice(x)
	for i := 0; i < len(nums) / 2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
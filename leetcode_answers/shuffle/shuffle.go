package main

import "fmt"

func main() {}

func shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)

	for i:=0; i < n*2; i=i+2 {
		fmt.Println(i)
		result[i] = nums[i/2]
		result[i+1] = nums[n+i/2]
	}
    return result
}

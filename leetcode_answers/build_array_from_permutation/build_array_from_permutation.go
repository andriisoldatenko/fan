package main

func main() {}

func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for index, num := range nums {
		result[index] = nums[num]
	}
	return result
}

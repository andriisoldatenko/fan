package main

func main() {}

func sum(nums []int) int {
	result := 0
	for _, el := range nums {
		result += el
	}
	return result
}

func pivotIndex(nums []int) int {
	left := 0
	right := sum(nums)

	for i := 0; i < len(nums); i++ {
		right = right - nums[i]
		if left == right {
			return i
		}
		left = left + nums[i]
	}
	return -1
}

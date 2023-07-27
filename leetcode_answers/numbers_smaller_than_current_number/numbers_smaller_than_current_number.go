package main

func main() {}

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] > nums[j] {
				count++
			}
		}
		result[i] = count
	}
	return result
}

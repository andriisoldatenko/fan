package main

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	end := 0
	for index, i := range nums {
		end += i
		result[index] = end
	}
	return result
}

func main() {
}

package main

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Sum[T Signed](s []T) T {
	var sum T
	for _, el := range s {
		sum += el
	}
	return sum
}

func Abs[T Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func leftRightDifference(nums []int) []int {
	left := 0
	right := Sum(nums)
	answer := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		right = right - nums[i]
		answer[i] = Abs(left - right)
		left = left + nums[i]
	}
	return answer
}

func main() {
	leftRightDifference([]int{10, 4, 8, 3})
}

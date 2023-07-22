package main

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Max[T Signed](s []T) T {
	max := s[0]
	for _, arg := range s[1:] {
		if arg > max {
			max = arg
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := Max(candies)
	result := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= max {
			result[i] = true
		}
	}
	return result
}

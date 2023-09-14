package main

func main() {}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, row := range accounts {
		localSum := 0
		for _, j := range row {
			localSum += j
		}
		max = Max(localSum, max)
	}
	return max
}

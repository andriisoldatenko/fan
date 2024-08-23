package main

import (
	"math/bits"
)

func main() {}

func findComplement(num int) int {
	x := uint(num)
	b := bits.Len(x)
	// bits.Len returns 0 when input is 0 (but we want start from 1)
	if b == 0 {
		b = 1
	}
	max_ := (2 << (b - 1)) - 1
	return max_ - num
}

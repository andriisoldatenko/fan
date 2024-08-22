package main

import (
	"fmt"
)

// The following Go function will compute the Hamming distance of two integers
// (considered as binary values, that is, as sequences of bits). The running
// time of this procedure is proportional to the Hamming distance rather
// than to the number of bits in the inputs. It computes the bitwise exclusive
// or of the two inputs, and then finds the Hamming weight of the result
// (the number of nonzero bits) using an algorithm of Wegner (1960) that
// repeatedly finds and clears the lowest-order nonzero bit.
func hammingDistance(x int, y int) int {
	dist := 0
	// bitwise XOR
	val := x ^ y
	for val != 0 {
		// A bit is set, so increment the count and clear the bit
		dist++
		val &= val - 1
	}
	return dist
}

func main() {
	fmt.Println(hammingDistance(252222, 16))
}

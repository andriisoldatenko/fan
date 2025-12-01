package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var maxAttempts int = 2

	var initialDelay = time.Second / 2

	var totalDelay time.Duration

	for i := range maxAttempts {
		delay := initialDelay * time.Duration(math.Pow(2, float64(i)))
		totalDelay += delay
	}

	fmt.Println(totalDelay.String())
	fmt.Println(1 / math.Pow(2, float64(maxAttempts)))
}

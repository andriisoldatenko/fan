package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)
	if (x - 2) > 0 && (x - 2) % 2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
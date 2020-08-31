package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)
	if x >= 30 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
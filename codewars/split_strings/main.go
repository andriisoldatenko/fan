package main

import "fmt"

func Solution(str string) []string {
	newStr := str
	if len(str)%2 != 0 {
		newStr = fmt.Sprintf("%s_", str)
	}
	res := []string{}
	for i := 0; i < len(str); i = i + 2 {
		res = append(res, fmt.Sprintf("%s%s", string(newStr[i]), string(newStr[i+1])))
	}
	return res
}

func main() {}

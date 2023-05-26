package main

func main() {}

func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, op := range operations {
		if op == "++X" || op == "X++" {
			result++
		} else {
			result--
		}
	}
	return result
}

package main

func main() {}

func decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)

	result[0] = first
	for i := 0; i < len(encoded); i++ {
		first = first ^ encoded[i]
		result[i+1] = first
	}
	return result
}

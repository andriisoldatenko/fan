package main

func numJewelsInStones(J string, S string) int {
	total := 0
	for _, i := range J {
		for _, j := range S {
			if i == j {
				total += 1
			}
		}
	}
	return total
}

func main() {
	numJewelsInStones("z", "ZZ")
}

package main

func main() {}

var cache = make(map[int]int)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	result := fib(n-1) + fib(n-2)
	cache[n] = result
	return result
}

func fib1(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 1; i < n; i++ {
		a += b
		a, b = b, a
	}
	return b
}

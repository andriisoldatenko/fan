package main

func main() {}

func prefixProduct(a []int) []int {
	for i := 1; i < len(a); i++ {
		a[i] = a[i] * a[i-1]
	}
	return a
}

func suffixProduct(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = 1
	}

	for i := len(a) - 2; i >= 0; i-- {
		res[i] = res[i+1] * a[i+1]
	}
	return res
}

func productExceptSelf(nums []int) []int {
	//1, 2, 3, 4
	//
	//2*3*4; 3*4; 4;
	//
	//            1; 1*2; 1*2*3
	//
	//2*3*4 = 24
	//
	//3*4*1 = 12
	//
	//4*1*2 = 6
	//
	//1*2*3 = 6
	// remove first
	suffix := suffixProduct(nums)
	//fmt.Println(suffix)
	// remove last
	prefix := prefixProduct(nums)
	prefix = prefix[:len(prefix)-1]
	//fmt.Println(prefix)

	result := make([]int, len(nums))
	result[0] = suffix[0]
	for i := 0; i < len(prefix); i++ {
		result[i+1] = prefix[i] * suffix[i+1]
	}
	return result
	// You must write an algorithm that runs in O(n) time and without using the division operation.
}

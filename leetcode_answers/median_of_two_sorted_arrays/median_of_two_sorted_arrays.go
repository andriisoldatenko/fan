package main

func main() {}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := make([]int, len(nums1)+len(nums2))

	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result[k] = nums1[i]
			k++
			i++
		} else {
			result[k] = nums2[j]
			k++
			j++
		}
	}

	for i < len(nums1) {
		result[k] = nums1[i]
		k++
		i++
	}
	for j < len(nums2) {
		result[k] = nums2[j]
		k++
		j++
	}

	if len(result)%2 == 0 {
		index := len(result) / 2
		return float64(result[index-1]+result[index]) / 2.0
	} else {
		return float64(result[len(result)/2])
	}
}

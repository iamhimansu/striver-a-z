package main

import "fmt"

func unionArray(nums1 []int, nums2 []int) []int {
	length1 := len(nums1)
	length2 := len(nums2)

	union := make([]int, 0, length1+length2)

	// 3 4 6 7 9 9 10 || 1 5 7 8 8
	i, j := 0, 0
	for i < length1 && j < length2 {
		var next int
		if nums1[i] < nums2[j] {
			next = nums1[i]
			i++
		} else if nums2[j] < nums1[i] {
			next = nums2[j]
			j++
		} else { // Elements are equal
			next = nums1[i]
			i++
			j++
		}

		if len(union) == 0 || union[len(union)-1] != next {
			union = append(union, next)
		}
	}
	for ; i < length1; i++ {
		if len(union) == 0 || union[len(union)-1] != nums1[i] {
			union = append(union, nums1[i])
		}
	}

	for ; j < length2; j++ {
		if len(union) == 0 || union[len(union)-1] != nums2[j] {
			union = append(union, nums2[j])
		}
	}
	return union
}

func main() {
	nums1 := []int{3, 4, 6, 7, 9, 9, 10}
	nums2 := []int{1, 5, 7, 8, 8}
	union := unionArray(nums1, nums2)
	fmt.Printf("%v", union)
}

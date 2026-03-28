package main

import "fmt"

/** compare adjacent element and swap the largest to the right */
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	arr := []int{5, 4, 4, 1, 1}
	bubbleSortedArray := bubbleSort(arr)

	fmt.Printf("%v", bubbleSortedArray)
}

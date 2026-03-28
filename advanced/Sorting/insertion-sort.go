package main

import "fmt"

func insertionSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
	return nums
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	insertionSort(arr)
	fmt.Printf("%v", arr)
}

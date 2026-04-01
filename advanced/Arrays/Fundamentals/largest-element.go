package main

import "fmt"

func largestElement(nums []int) int {
	currentLargest := -1
	if len(nums) > 0 {
		currentLargest = nums[0]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > currentLargest {
			currentLargest = nums[i]
		}
	}

	return currentLargest
}

func main() {
	nums := []int{12, 35, 1, 10, 34, 33, 1}
	fmt.Println(largestElement(nums))
}

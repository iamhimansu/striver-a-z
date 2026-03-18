package main

import (
	"fmt"
)

func secondLargestElement(nums []int) int {

	if len(nums) < 2 {
		return -1
	}

	currentLargest := nums[0]
	secondLargest := nums[0]
	hasSecond := false

	for i := 1; i < len(nums); i++ {
		current := nums[i]

		if current > currentLargest {
			secondLargest = currentLargest
			currentLargest = current
			hasSecond = true
		} else if current != currentLargest {
			if !hasSecond || current > secondLargest {
				secondLargest = current
				hasSecond = true
			}
		}
	}

	if !hasSecond {
		return -1
	}

	return secondLargest
}

func main() {
	nums := []int{100, -100}
	fmt.Println(secondLargestElement(nums))
}

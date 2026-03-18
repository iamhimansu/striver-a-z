package main

import (
	"fmt"
	"math"
)

func secondLargestElement(nums []int) int {

	currentLargest := nums[0]
	secondLargest := math.MinInt

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		if current > currentLargest {
			secondLargest = currentLargest
			currentLargest = current
		} else if current > secondLargest && current != currentLargest {
			secondLargest = current
		}
	}

	// no second largest found
	if secondLargest == math.MinInt {
		return -1
	}

	return secondLargest
}

func main() {
	nums := []int{100, -100}
	fmt.Println(secondLargestElement(nums))
}

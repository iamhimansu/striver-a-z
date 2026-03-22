package main

import "fmt"

func init() {
	fmt.Println("Initialized")
}

func findMaxConsecutiveOnes(nums []int) int {
	currentCount := 0
	maxCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}

		} else {
			currentCount = 0
		}
	}
	return maxCount
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 0, 1}))
}

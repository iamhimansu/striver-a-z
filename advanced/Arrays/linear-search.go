package main

import "fmt"

func linearSearch(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}

	}
	return -1
}

func main() {
	nums := []int{0, 3, 5, -4, 1}
	targetIndex := linearSearch(nums, 1)
	fmt.Println(targetIndex)
}

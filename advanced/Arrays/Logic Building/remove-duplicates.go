package main

import "fmt"

func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func main() {
	arr := []int{0, 0, 3, 3, 5, 6}
	uniqueElements := removeDuplicates(arr)

	fmt.Printf("%v", uniqueElements)
}

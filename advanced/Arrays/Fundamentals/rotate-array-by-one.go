package main

import "fmt"

func rotateArrayByOne(nums []int) {
	if len(nums) > 0 {
		first := nums[0]
		copy(nums, nums[1:])
		nums[len(nums)-1] = first
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	rotateArrayByOne(nums)
	fmt.Println(nums)
}

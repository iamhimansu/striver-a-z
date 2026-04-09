package main

import "fmt"

func Sum(nums []int, max int) int {
	if max == 0 {
		return 0
	}
	return nums[max-1] + Sum(nums, max-1)
}
func arraySum(nums []int) int {
	return Sum(nums, len(nums))
}

func main() {
	arr := []int{1, 4, 5, 6}
	sum := arraySum(arr)

	fmt.Printf("%v", sum)
}

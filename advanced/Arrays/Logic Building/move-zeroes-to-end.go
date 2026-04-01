package main

import "fmt"

func moveZeroesToEnd(nums []int) []int {
	size := len(nums)
	j := size - 1
	for i := 0; i < size; i++ {
		if nums[i] == 0 && i < j {
			//swap last with zero
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return nums
}

func main() {

	arr := []int{0, 0, 0, 1, 3, -2}
	mzte := moveZeroesToEnd(arr)

	fmt.Printf("%v", mzte)
}

package main

import "fmt"

func moveZeroesToEndWithoutOrder(nums []int) []int {
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

func moveZeroesToEnd(nums []int) []int {
	/** loop and find the 1st non-zero
	swap it with zero position and swap the zero with non-zero position
	**/
	for i, j := 0, 0; i < len(nums); i++ {
		if j <= i && nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	return nums
}

func main() {

	arr := []int{1, -6, -9, 8, -9, 0, 0, 1, 3, -2}
	mztewo := moveZeroesToEndWithoutOrder(arr)
	mzte := moveZeroesToEnd(arr)

	fmt.Printf("%v", mztewo)
	fmt.Printf("%v", mzte)

}

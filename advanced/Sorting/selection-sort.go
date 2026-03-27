package main

import "fmt"

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}

	return nums
}

func main() {

	arr := []int{7, 4, 1, 5, 3}
	sorted := selectionSort(arr)

	fmt.Printf("%v", arr)
	fmt.Print(sorted)

}

package main

import "fmt"

func largestElement(nums []int) int {

	currentLargest := -1
	current := 0
	next := current

	//[12, 35, 1, 10, 34, 1]
	for i := 0; i < len(nums)-1; i++ {

		current = nums[i]
		next = nums[i+1]

		//next largest
		if next > currentLargest {
			currentLargest = next
		}

	}

	return currentLargest

}

func main() {
	nums := []int{12, 35, 1, 10, 34, 33, 1}
	fmt.Println(largestElement(nums))
}

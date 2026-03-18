package main

import "fmt"

func secondLargestElement(nums []int) int {

	currentLargest := -1
	current := 0
	next := current

	//[12, 35, 1, 10, 34, 1]
	for i := 0; i < len(nums)-1; i++ {

		current = nums[i]
		next = nums[i+1]

		//next largest
		if current < next && next > currentLargest {
			currentLargest = next
		}

	}

	if currentLargest > 0 {
		return currentLargest
	}

	return -1
}

func main() {
	nums := []int{12, 35, 1, 10, 34, 1}
	fmt.Println(secondLargestElement(nums))
}

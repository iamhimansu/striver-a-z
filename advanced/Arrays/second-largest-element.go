package main

import "fmt"

func secondLargestElement(nums []int) int {

	currentLargest := 0
	current := 0
	secondLargest := -1

	//[12, 35, 1, 10, 34, 33, 1]
	for i := 0; i < len(nums)-1; i++ {

		current = nums[i]

		//next largest
		if current > currentLargest {
			secondLargest = currentLargest
			currentLargest = current
		} else if current > secondLargest && current != currentLargest {
			secondLargest = current
		}

	}

	return secondLargest

}

func main() {
	nums := []int{12, 35, 1, 32, 10, 30, 31, 1}
	fmt.Println(secondLargestElement(nums))
}

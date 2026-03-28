package main

import "fmt"

func mergeSort(nums []int) []int {
	mergeSortHelper(nums, 0, len(nums)-1)
	return nums
}
func mergeSortHelper(nums []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSortHelper(nums, low, mid)
	mergeSortHelper(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []int, low int, mid int, high int) {
	left, right := low, mid+1
	temp := make([]int, 0, high-low+1)

	for left <= mid && right <= high {
		if nums[left] < nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, nums[left])
		left++
	}

	for right <= high {
		temp = append(temp, nums[right])
		right++
	}

	for i := low; i <= high; i++ {
		nums[i] = temp[i-low]
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1, 9, 7, 8, 0}
	sortedArray := mergeSort(arr)

	fmt.Printf("%v", sortedArray)
}

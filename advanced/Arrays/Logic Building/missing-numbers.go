package main

import "fmt"

func findMissingNumbers(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2
	fmt.Println(sum)
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}

func main() {
	arr := []int{0, 2, 3, 1, 4}
	missing := findMissingNumbers(arr)
	fmt.Printf("%d", missing)
}

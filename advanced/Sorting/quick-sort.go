package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	i, j := low, high
	for i < j {
		//[i=4, 5, 3, 2, 7, j=1]
		//[1, 2, 3, 4, 5, 7]
		//4 < 5
		// left side
		for arr[i] <= pivot && i <= high {
			i++
		}
		//right side
		//4>1
		for arr[j] > pivot && j > low {
			j--
		}
		if i < j {
			//swap
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

func quickSort(arr []int, low int, high int) []int {
	if low >= high {
		return arr
	}
	pIndex := partition(arr, low, high)

	quickSort(arr, low, pIndex-1)
	quickSort(arr, pIndex+1, high)
	return arr
}

func main() {
	arr := []int{4, 6, 2, 5, 7, 9, 1, 3}
	sorted := quickSort(arr, 0, len(arr)-1)
	fmt.Printf("%v", sorted)
}

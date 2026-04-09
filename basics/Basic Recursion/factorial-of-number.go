package main

import "fmt"

func factorial(n int) int64 {
	if n <= 0 {
		return 1
	}
	return int64(n) * factorial(n-1)
}

func main() {
	fact := factorial(5)

	fmt.Printf("%v", fact)
}

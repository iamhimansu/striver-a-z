package main

import "fmt"

func sumTillN(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sumTillN(n-1)
}

func main() {
	n := sumTillN(4)

	fmt.Printf("%v", n)

}

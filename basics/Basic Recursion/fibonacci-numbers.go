package main

import "fmt"

func printFib(a, b, limit int) {
	if a > limit {
		return
	}
	fmt.Printf("%d ", a)
	printFib(b, a+b, limit) // Pass the next two numbers forward
}

func main() {
	printFib(0, 1, 50)
}

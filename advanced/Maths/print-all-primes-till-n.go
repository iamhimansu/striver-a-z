package main

import (
	"fmt"
	"slices"
)

func primeTillN(n int) []int {
	allPrimes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		allPrimes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		for j := i * i; j <= n; j += i {
			if allPrimes[j] {
				allPrimes[j] = false
			}
		}

	}

	primes := []int{}
	for i := 0; i <= n; i++ {
		if allPrimes[i] {
			primes = append(primes, i)
		}
	}

	slices.Sort(primes)

	fmt.Println("%v", primes)
	return primes
}

func main() {

	primeTillN(300)
}

package main

import "fmt"

func rev(s []rune, i int, j int) []rune {
	if i > j {
		return s
	}
	s[i], s[j] = s[j], s[i]
	rev(s, i+1, j-1)

	return s
}
func reverseString(s []rune) []rune {
	return rev(s, 0, len(s)-1)
}

func main() {
	s := []rune{'h', 'e', 'l', 'l', 'o'}

	fmt.Printf("%c", reverseString(s))
}

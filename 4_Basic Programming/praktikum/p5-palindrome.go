package main

import (
	"fmt"
	"os"
)

func checkIsPalindrome(clause string) bool {
	for i := 0; i < len(clause)/2; i++ {
		if clause[i] != clause[len(clause)-(i+1)] {
			return false
		}
	}

	return true
}

func main() {
	clause := os.Args[1]
	fmt.Println("Palindrome: ", checkIsPalindrome(clause))
}

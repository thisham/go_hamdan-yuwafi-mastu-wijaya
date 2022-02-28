package main

import (
	"fmt"
	"strings"
)

func getDisplayedOnceNumbers(numbers string) []string {
	var results []string

	// using O(n)
	for i := 0; i < len(numbers); i++ {
		if strings.Count(numbers, string(numbers[i])) == 1 {
			results = append(results, string(numbers[i]))
		}
	}

	return results
}

func main() {
	fmt.Println(getDisplayedOnceNumbers("6787871234123"))
}

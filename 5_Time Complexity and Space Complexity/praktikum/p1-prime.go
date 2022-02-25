package main

import (
	"fmt"
)

func checkIsPrimeNumber(sourceNumber int) bool {
	var divideFactors []int

	// still using O(n) complexity
	for i := 1; i <= sourceNumber; i++ {
		if sourceNumber%i == 0 {
			divideFactors = append(divideFactors, i)
		}
	}

	return len(divideFactors) == 2
}

func main() {
	fmt.Println(checkIsPrimeNumber(13))
	fmt.Println(checkIsPrimeNumber(15))
}

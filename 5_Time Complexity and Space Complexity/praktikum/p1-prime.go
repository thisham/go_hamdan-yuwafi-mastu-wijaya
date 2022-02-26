package main

import (
	"fmt"
	"math"
)

func checkIsPrimeNumber(sourceNumber int) bool {
	absoluteSourceNumber := int(math.Abs(float64(sourceNumber)))
	if absoluteSourceNumber < 2 {
		return false
	}

	if absoluteSourceNumber%2 == 0 {
		return absoluteSourceNumber == 2
	}

	root := int(math.Sqrt(float64(absoluteSourceNumber)))

	for i := 3; i <= root; i += 2 {
		if sourceNumber%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkIsPrimeNumber(13))
	fmt.Println(checkIsPrimeNumber(15))
}

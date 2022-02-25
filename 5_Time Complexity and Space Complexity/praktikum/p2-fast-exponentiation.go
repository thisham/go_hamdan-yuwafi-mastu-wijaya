package main

import "fmt"

func powerify(base int, power int) int {
	var result int

	// using O(n)
	for i := 0; i < power; i++ {
		if result == 0 {
			result = base
			continue
		}
		result = result * base
	}

	return result
}

func main() {
	fmt.Println(powerify(7, 3))
}

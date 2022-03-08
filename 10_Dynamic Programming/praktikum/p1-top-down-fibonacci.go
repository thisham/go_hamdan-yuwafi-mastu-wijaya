package main

import (
	"fmt"
	"os"
	"strconv"
)

func findFibonacciAt(index int) int {
	if index <= 1 {
		return index
	}

	const maxLookup = 100
	var fibonacci = make([]int, maxLookup)

	fibonacci[index] = findFibonacciAt(index-1) + findFibonacciAt(index-2)
	return fibonacci[index]
}

func main() {
	var err error
	var arguments = os.Args[1:]
	nums := make([]int, len(arguments))

	for index, num := range nums {
		if num, err = strconv.Atoi(arguments[index]); err != nil {
			panic(err)
		}
		fmt.Println("index:", num, "fibonacci:", findFibonacciAt(num))
	}
}

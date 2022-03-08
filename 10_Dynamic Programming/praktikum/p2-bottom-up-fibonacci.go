package main

import (
	"fmt"
	"os"
	"strconv"
)

func findFibonacciAt(index int) int {
	if index <= 0 {
		return 0
	}

	var fibonacci = make([]int, index+1)
	fibonacci[0], fibonacci[1] = 0, 1

	for i := 2; i <= index; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
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

package main

import "fmt"

func getFibonacci(position int) int {
	if position <= 0 {
		return 0
	}

	if position == 1 {
		return 1
	}

	number := getFibonacci(position-1) + getFibonacci(position-2)
	return number
}

func main() {
	fmt.Println(getFibonacci(0))
	fmt.Println(getFibonacci(2))
	fmt.Println(getFibonacci(9))
}

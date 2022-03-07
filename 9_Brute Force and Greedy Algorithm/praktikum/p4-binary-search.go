package main

import "fmt"

func search(source []int, query int) (result int, steps int, index int) {
	var left int = 0
	var right int = len(source) - 1
	index = -1

	for (left <= right) && index <= -1 {
		center := (left + right) / 2
		if query < source[center] {
			right = center - 1
		} else if query > source[center] {
			left = center + 1
		} else {
			index = center
		}
		steps++
	}

	if index <= -1 {
		fmt.Println("search query not found in array source! completely searching in", steps, "steps.")
		return
	}

	return source[index], steps, index
}

func main() {
	fmt.Println(search([]int{1, 1, 3, 5, 5, 6, 7}, 3)) // 3 3 2
	fmt.Println(search([]int{1, 1, 3, 5, 5, 6, 7}, 9)) // panic
}

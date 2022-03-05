package main

import "fmt"

type BorderType struct {
	value int
	index int
}

type Border struct {
	min BorderType
	max BorderType
}

func findBorderAtIndex(source []int) string {
	var result Border

	result.min.index = 0
	result.min.value = source[result.min.index]
	result.max.index = 0
	result.max.value = source[result.max.index]

	for index, number := range source {
		if result.min.value > number {
			result.min.value = number
			result.min.index = index
		}
		if result.max.value < number {
			result.max.value = number
			result.max.index = index
		}
	}

	return fmt.Sprint("min: ", result.min.value, " at index ", result.min.index, "; max:", result.max.value, " at index ", result.min.index)
}

func main() {
	fmt.Println(findBorderAtIndex([]int{5, 7, 4, -2, -1, 8}))
}

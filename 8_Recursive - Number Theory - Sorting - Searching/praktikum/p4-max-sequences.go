package main

import "fmt"

func findSummary(things []int) (result int) {
	for _, thing := range things {
		result += thing
	}
	return
}

func findMaxSequenceSums(source []int) int {
	var max, current, index int

	for len(source) > 0 {
		if index == 0 {
			current = findSummary(source)
			max = current
			index++
			continue
		}

		if index%2 != 0 {
			source = source[1:]
			current = findSummary(source)
			if max < current {
				max = current
			}
		}

		if index%2 == 0 {
			source = source[0 : len(source)-1]
			current = findSummary(source)
			if max < current {
				max = current
			}
		}
		index++
	}
	return max
}

func main() {
	fmt.Println(findMaxSequenceSums([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(findMaxSequenceSums([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
}

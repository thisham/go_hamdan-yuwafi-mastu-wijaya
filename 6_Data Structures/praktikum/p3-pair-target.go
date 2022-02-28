package main

import "fmt"

func findPairs(source []int, target int) (int, int) {
	if len(source) > 1 {

		temp := make(map[int]int)

		for i, value := range source {
			j, numberExists := temp[target-value]
			if numberExists {
				return source[j], source[i]
			}

			temp[source[i]] = i
		}
	}
	return -1, -1
}

func main() {
	fmt.Println(findPairs([]int{1, 4, 1, 4, 7, 5}, 9))
}

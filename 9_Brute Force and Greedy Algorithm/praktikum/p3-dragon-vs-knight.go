package main

import (
	"fmt"
	"sort"
)

func sortAscending(source []int) (result []int) {
	sort.Slice(source, func(i, j int) bool {
		return source[i] < source[j]
	})
	return source
}

// This algorithm eliminates unable dragon slashers depends on
// knight's height. When the knight is not able to slash a dragon,
// he falls down. And, each knight is tackle a dragon ahead.
func getMinimumOfKnightsTotalHeight(dragonHeads, knightHeights []int) int {
	var knightsMinimumHeight int = 0

	// sort
	dragonHeads = sortAscending(dragonHeads)
	knightHeights = sortAscending(knightHeights)

	for _, dragonHead := range dragonHeads {
		knights := len(knightHeights)

		// search for the knight that able to slash the dragon
		if searchIndex := sort.Search(knights, func(i int) bool { return knightHeights[i] >= dragonHead }); searchIndex < knights {
			// when found, add to summary
			knightsMinimumHeight += knightHeights[searchIndex]
			knightHeights = knightHeights[searchIndex+1:]
			continue
		}

		defer fmt.Println("knight falls.")
		return -1
	}
	return knightsMinimumHeight
}

func main() {
	fmt.Println(getMinimumOfKnightsTotalHeight([]int{5, 4}, []int{7, 8, 4}))    // 11
	fmt.Println(getMinimumOfKnightsTotalHeight([]int{5, 10}, []int{5}))         // -1 defer "knight falls"
	fmt.Println(getMinimumOfKnightsTotalHeight([]int{7, 2}, []int{4, 3, 1, 2})) // -1 defer "knight falls"
	fmt.Println(getMinimumOfKnightsTotalHeight([]int{7, 2}, []int{2, 1, 8, 5})) // 10
}

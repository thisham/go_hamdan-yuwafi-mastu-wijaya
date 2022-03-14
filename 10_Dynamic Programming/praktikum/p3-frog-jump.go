package main

import (
	"fmt"
	"math"
)

// type stone map[int]bool
// type any map[pair]

// func canCross(map[int]*stones) {

// }

func jump(stones []int) int {
	var initialStone int = 0
	var numberOfStones int = len(stones)

	// fill initial jumps
	var costs = make([]int, numberOfStones)
	for key := range costs {
		costs[key] = -1
	}

	// initial stone doesn't had a cost to jump
	costs[initialStone] = 0

	for i := 1; i < numberOfStones; i++ {
		var jumpTwice int = math.MaxInt
		var jumpOnce int = costs[i-1] + int(math.Abs(float64(stones[i]-stones[i-1])))
		if i > 1 {
			jumpTwice = costs[i-2] + int(math.Abs(float64(stones[i]-stones[i-2])))
		}

		costs[i] = int(math.Min(float64(jumpOnce), float64(jumpTwice)))
	}

	return costs[len(costs)-1]
}

func main() {
	fmt.Println(jump([]int{10, 30, 40, 20}))         //30
	fmt.Println(jump([]int{30, 10, 60, 10, 60, 50})) //40
}

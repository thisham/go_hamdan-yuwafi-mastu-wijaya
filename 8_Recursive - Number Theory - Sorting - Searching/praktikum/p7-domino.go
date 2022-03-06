package main

import "fmt"

func playDominosAcceptable(cards [][]int, deck [2]int) []int {
	// swap deck
	if deck[1] > deck[0] {
		deck[0], deck[1] = deck[1], deck[0]
	}

	// find pairs
	for _, card := range cards {
		if card[0] == deck[0] || card[1] == deck[0] {
			return card
		}
	}

	// exit when no one matched
	panic("tutup kartu")
}

func main() {
	fmt.Println(playDominosAcceptable([][]int{[]int{6, 4}, []int{5, 2}, []int{3, 2}, []int{1, 2}}, [2]int{2, 4}))
}

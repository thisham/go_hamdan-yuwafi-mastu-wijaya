package main

import (
	"fmt"
	"sort"
)

func sortAscendingPrices(productPrices []int) []int {
	sort.Slice(productPrices, func(i, j int) bool {
		return productPrices[i] < productPrices[j]
	})
	return productPrices
}

func countProductCanBeBuyed(money int, productPrices []int) int {
	var result int

	for _, productPrice := range productPrices {
		if money < productPrice {
			break
		}

		money -= productPrice
		result++
	}
	return result
}

// sort dulu
// gunakan greedy alg.
// return as int, berapa barag yang dibeli
func getMaximumBuyProduct(money int, productPrices []int) int {
	var sortedDescendingPrices = sortAscendingPrices(productPrices)
	var thingsCanBeBuyed = countProductCanBeBuyed(money, sortedDescendingPrices)
	return thingsCanBeBuyed
}

func main() {
	fmt.Println("things can be buyed:", getMaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000, 5000}))
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func getCoins(money int) []int {
	var shards = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	var results []int

	for money > 0 {
		for _, coin := range shards {
			if coin > money {
				continue
			}
			money -= coin
			results = append(results, coin)
			break
		}
	}

	return results
}

func main() {
	var err error
	var arguments = os.Args[1:]
	nums := make([]int, len(arguments))

	for index, num := range nums {
		if num, err = strconv.Atoi(arguments[index]); err != nil {
			panic(err)
		}
		fmt.Println("money:", num, "result:", getCoins(num))
	}
}

package main

import (
	"fmt"
)

func getCommonSubstring(base, pair string) string {
	var result string

	if len(base) < len(pair) {
		base, pair = pair, base
	}

	var i int = 0

	for i < len(base) {
		var j int = 0

		if base[i] == pair[j] {
			for j <= len(pair) {
				if base[i:i+j] == pair[0:j] {
					result = pair[0:j]
					j++
				} else {
					i += j
					break
				}
			}
		}
		i++
	}

	return result
}

func main() {
	fmt.Println(getCommonSubstring("kangaroo", "kang"))
	fmt.Println(getCommonSubstring("ast", "middleeast"))
}

package main

import "fmt"

func mergeArray(base, target []string) []string {
	base = append(base, target...)
	check := make(map[string]int)
	result := make([]string, 0)

	for _, value := range base {
		check[value] = 1
	}

	for character := range check {
		result = append(result, character)
	}

	return result
}

func main() {
	fmt.Println(mergeArray([]string{"king", "save", "queen"}, []string{"but", "king", "was", "angry"}))
}

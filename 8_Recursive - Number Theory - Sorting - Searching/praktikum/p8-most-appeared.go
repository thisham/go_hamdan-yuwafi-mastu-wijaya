package main

import "fmt"

type Pair struct {
	name  string
	count int
}

type KeyValuePair map[string]int

func groupItemsByName(items []string) KeyValuePair {
	var result = make(KeyValuePair)
	for _, item := range items {
		result[item] += 1
	}
	return result
}

func pairValueInStruct(groupedValue KeyValuePair) (result []Pair) {
	for key, value := range groupedValue {
		result = append(result, Pair{name: key, count: value})
	}
	return
}

func findMostAppearedValue(items []string) []Pair {
	var groupedValuesWithKeys = make(KeyValuePair)
	groupedValuesWithKeys = groupItemsByName(items)
	var pairs = pairValueInStruct(groupedValuesWithKeys)
	return pairs
}

func main() {
	fmt.Println(findMostAppearedValue([]string{"js", "js", "golang", "ruby", "java", "rust", "rust", "rust"}))
}

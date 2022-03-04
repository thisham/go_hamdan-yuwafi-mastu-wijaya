package main

import "fmt"

func getNumberAtIndex(position int) (number int) {
	// if position <= 1 {
	// 	return 2
	// }

	// number = getNumberAtIndex(position - 1)
	// number += 1

	number = 0
	currentPosition := 0
	currentCountedNumber := 0

	for currentPosition <= position {
		switch {
		case bool(currentCountedNumber != 2 && currentCountedNumber%2 == 0):
			currentCountedNumber++
			continue
		case bool(currentCountedNumber != 3 && currentCountedNumber%3 == 0):
			currentCountedNumber++
			continue
		case bool(currentCountedNumber != 5 && currentCountedNumber%5 == 0):
			currentCountedNumber++
			continue
		case bool(currentCountedNumber != 7 && currentCountedNumber%7 == 0):
			currentCountedNumber++
			continue
		default:
			number = currentCountedNumber
			currentPosition++
			currentCountedNumber++
		}
	}

	return number
}

func main() {
	fmt.Println(getNumberAtIndex(-1))
	fmt.Println(getNumberAtIndex(1))
	fmt.Println(getNumberAtIndex(2))
	fmt.Println(getNumberAtIndex(10))
}

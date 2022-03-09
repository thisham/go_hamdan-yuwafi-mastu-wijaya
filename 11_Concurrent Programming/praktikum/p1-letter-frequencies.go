package main

import (
	"fmt"
	"os"
	"strings"
)

type LetterFrequency map[rune]int

func countLetterFrequency(word string) LetterFrequency {
	var frequency = LetterFrequency{}
	for _, character := range word {
		frequency[character]++
	}
	return frequency
}

func main() {
	var sentence = os.Args[1]
	var words = strings.Split(sentence, " ")

	result := LetterFrequency{}
	resultChannel := make(chan LetterFrequency, len(words))

	// concurrently count the letters of the words
	for _, word := range words {
		go func(word string) {
			resultChannel <- countLetterFrequency(word)
		}(word)
	}

	// totaling each concurrent
	for i := 0; i < len(words); i++ {
		for letter, frequency := range <-resultChannel {
			result[letter] += frequency
		}
	}

	// output
	for letter, count := range result {
		fmt.Println(string(letter), count)
	}
}

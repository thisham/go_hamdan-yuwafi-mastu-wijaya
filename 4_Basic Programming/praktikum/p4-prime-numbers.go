package main

import "fmt"

func detectPrimeNumber(sourceNumber int) bool {
	var divideFactors []int

	for i := 1; i <= sourceNumber; i++ {
		if sourceNumber%i == 0 {
			divideFactors = append(divideFactors, i)
		}
	}

	return len(divideFactors) == 2
}

func main() {
	var sourceNumber int
	var isPrimeNumber bool

	fmt.Printf("Masukkan suatu bilangan: ")
	fmt.Scanf("%d", &sourceNumber)
	isPrimeNumber = detectPrimeNumber(sourceNumber)

	if isPrimeNumber {
		fmt.Println("Bilangan prima")
	} else {
		fmt.Println("Bukan bilangan prima")
	}
}

package main

import "fmt"

func findDivideFactor(sourceNumber int) []int {
	var divideFactors []int

	for i := 1; i <= sourceNumber; i++ {
		if sourceNumber%i == 0 {
			divideFactors = append(divideFactors, i)
		}
	}

	return divideFactors
}

func main() {
	var number int

	fmt.Printf("Masukkan suatu angka: ")
	fmt.Scanf("%d", &number)
	fmt.Println("Faktor pembagi: ", findDivideFactor(number))
}

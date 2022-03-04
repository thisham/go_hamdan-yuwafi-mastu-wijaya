package main

import "fmt"

func getBunchOfPrime(starting, amount int) []int {
	counting := 0
	var primeNumbers []int

	for counting < amount {
		switch {
		case bool(starting != 2 && starting%2 == 0):
			break
		case bool(starting != 3 && starting%3 == 0):
			break
		case bool(starting != 5 && starting%5 == 0):
			break
		case bool(starting != 7 && starting%7 == 0):
			break
		default:
			primeNumbers = append(primeNumbers, starting)
			counting++
		}
		starting++
	}
	return primeNumbers
}

func squarePrime(height, width, starting int) {
	if height < 1 || width < 1 {
		panic("invalid params caught!")
	}

	var primeNumbers []int = getBunchOfPrime(starting, height*width)

	primeNumberIndex := 0
	primeNumbersSum := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("%d\t", primeNumbers[primeNumberIndex])
			primeNumbersSum += primeNumbers[primeNumberIndex]
			primeNumberIndex++
		}
		fmt.Println("")
	}

	fmt.Println("prime amount:", primeNumbersSum)
}

func main() {
	squarePrime(2, 3, 56)
}

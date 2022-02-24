package main

import "fmt"

func createMultiplicationTable(lastColumn int) {
	for i := 1; i <= lastColumn; i++ {
		for j := 1; j <= lastColumn; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Printf("\n")
	}
}

func main() {
	createMultiplicationTable(9)
}

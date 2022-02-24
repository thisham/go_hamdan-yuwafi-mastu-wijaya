package main

import "fmt"

func createAsteriskPyramid(base int) {
	for i := 0; i < base; i++ {
		for j := i; j < base; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func main() {
	createAsteriskPyramid(5)
}

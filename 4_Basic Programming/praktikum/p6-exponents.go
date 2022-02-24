package main

import (
	"fmt"
)

func getExponent(base int, power int) int {
	var result int

	for i := 0; i < power; i++ {
		if result == 0 {
			result = base
			continue
		}
		result = result * base
	}

	return result
}

func main() {
	var base, power int
	fmt.Printf("Masukkan angka dasar: ")
	fmt.Scanf("%d\n", &base)
	fmt.Printf("Masukkan angka pangkat: ")
	fmt.Scanf("%d\n", &power)

	fmt.Println("Hasil perpangkatan: ", getExponent(base, power))
}

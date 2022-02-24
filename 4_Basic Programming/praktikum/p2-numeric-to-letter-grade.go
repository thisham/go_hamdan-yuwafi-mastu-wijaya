package main

import "fmt"

func setNumericGrade() int {
	var numericValue int

	fmt.Printf("Masukkan nilai angka: ")
	fmt.Scanf("%d\n", &numericValue)
	return numericValue
}

func convertToLetter(numericValue int) string {
	if numericValue >= 80 && numericValue <= 100 {
		return "A"
	}
	if numericValue >= 65 && numericValue <= 79 {
		return "B"
	}
	if numericValue >= 50 && numericValue <= 64 {
		return "C"
	}
	if numericValue >= 35 && numericValue <= 49 {
		return "D"
	}
	if numericValue >= 0 && numericValue <= 34 {
		return "E"
	}
	return "invalid"
}

func main() {
	numericGrade := setNumericGrade()
	letterGrade := convertToLetter(numericGrade)
	fmt.Println("Hasil konversi nilai ke letter: ", letterGrade)
}

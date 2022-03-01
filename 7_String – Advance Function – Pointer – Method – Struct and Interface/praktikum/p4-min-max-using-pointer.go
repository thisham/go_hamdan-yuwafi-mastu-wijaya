package main

import "fmt"

func getBorder(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, number := range numbers {
		if min > *number {
			min = *number
		}
		if max < *number {
			max = *number
		}
	}

	return min, max
}

func main() {
	var a, b, c, d, e, f, min, max int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Scan(&e)
	fmt.Scan(&f)

	min, max = getBorder(&a, &b, &c, &d, &e, &f)
	fmt.Println("min", min, ", max", max)
}

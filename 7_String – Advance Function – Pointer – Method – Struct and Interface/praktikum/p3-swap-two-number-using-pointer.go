package main

import "fmt"

func swap(base, pair *int) {
	*base, *pair = *pair, *base
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}

package main

import "fmt"

func getRootVariant(a, b, c int) (x, y, z int) {
	for x <= a || x <= b || x <= c {
		y = 0
		for y <= a || y <= b || y <= c {
			z = 0
			for z <= a || z <= b || z <= c {
				isMatchOnEquationA := x+y+z == a
				isMatchOnEquationB := x*y*z == b
				isMatchOnEquationC := x*x+y*y+z*z == c
				if isMatchOnEquationA && isMatchOnEquationB && isMatchOnEquationC {
					return
				}
				z++
			}
			y++
		}
		x++
	}

	panic("not solved")
}

func main() {
	fmt.Println(getRootVariant(6, 6, 14)) // 1 2 3
	fmt.Println(getRootVariant(1, 2, 3))  // panic
}

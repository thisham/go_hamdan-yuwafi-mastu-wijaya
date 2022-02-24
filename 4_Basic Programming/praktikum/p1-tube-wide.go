package main

import (
	"fmt"
	"math"
)

type Tube struct {
	height float32
	radius float32
}

func (tube *Tube) setHeight() *Tube {
	fmt.Printf("Masukkan tinggi tabung: ")
	fmt.Scanf("%f\n", &tube.height)
	return tube
}

func (tube *Tube) setRadius() *Tube {
	fmt.Printf("Masukkan radius alas tabung: ")
	fmt.Scanf("%f\n", &tube.radius)
	return tube
}

func (tube *Tube) countWide() {
	tubeWide := 2 * math.Pi * tube.radius * (tube.radius + tube.height)
	fmt.Println("Luas tabung adalah", tubeWide)
}

func main() {
	tubeData := Tube{}
	tubeData.setHeight().setRadius().countWide()
}

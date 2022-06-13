package main

import (
	"fmt"
	"math"
)

func newton() float64 {
	r := 1000.0
	a := 2.0

	for i := 1; i <= 6; i++ {
		zahyou_x := a
		zahyou_y := 1.0
		for i := 0; i < 10; i++ {
			zahyou_y *= a
		}

		sessen_a := 10.0
		for i := 0; i < 9; i++ {
			sessen_a *= a
		}
		sessen_b := zahyou_y - sessen_a * zahyou_x

		next_a := (r - sessen_b) / sessen_a

		a = next_a
	}

	return a
}

func main(){
	expect := math.Pow(10, 0.3)
	fmt.Printf("Expected: %.15f\n", expect)

	ans := newton()
	fmt.Printf("Answer: %.15f\n", ans)
}

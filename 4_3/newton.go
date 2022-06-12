package main

import "fmt"

func main() {
	r := 2.0
	a := 2.0

	for i := 1; i <= 6; i++ {
		zahyou_x := a
		zahyou_y := a * a

		sessen_a := 2 * a
		sessen_b := zahyou_y - sessen_a * zahyou_x

		next_a := (r - sessen_b) / sessen_a

		fmt.Printf("Step %d %.15f -> %.15f\n", i, a, next_a)

		a = next_a
	}
}

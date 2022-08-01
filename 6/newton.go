package main

import (
	"fmt"
	"math"
)

func main() {
	// y = log2の近似値を調べる(自然対数)
	r := 2.0
	a := 2.0

	for i := 0; i < 7; i++ {
		zahyoX := a
		zahyoY := math.Exp(a)

		sessenA := math.Exp(a)
		sessenB := zahyoY - sessenA * zahyoX

		nextA := (r - sessenB) / sessenA

		fmt.Printf("Step %d = %.15f -> %.15f\n", i+1, a, nextA)
		a = nextA
	}
	fmt.Println(1.0/3.0)
}

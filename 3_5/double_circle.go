package main

import (
	"fmt"
	"math"
	"math/rand"
)

func dist(x, y float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2)
}

func main(){
	n := math.Pow(10, 6)
	m := 0.0
	ax, ay := 3.0, 3.0
	bx, by := 3.0, 7.0
	for i := 0; i <= int(n); i++{
		nx := rand.Float64() * 6
		ny := rand.Float64() * 9

		if dist(math.Abs(nx-ax), math.Abs(ny-ay)) <= 9{
			m += 1.0
			continue
		}

		if dist(math.Abs(nx-bx), math.Abs(ny-by)) <= 4 {
			m += 1.0
		}
	}
	fmt.Println(int(m))

	square := 6.0 * 9.0
	area := square * (m / n)
	fmt.Println(area)
}

package main

import (
	"fmt"
	"math"
	"math/rand"
)


func main() {
	n := math.Pow(10, 7)
	M := 0.0
	for i := 0; i < int(n); i++ {
		px := rand.Float64()
		py := rand.Float64()
		if (math.Pow(px, 2) + math.Pow(py, 2)) <= 1.0 {
			M += 1.0
		}
	}
	fmt.Println((M / n) * 4)
}

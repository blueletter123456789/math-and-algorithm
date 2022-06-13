package main

import "fmt"

func newton(){
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

func binSearch(){
	
	l, r := 1.0, 2.0
	sm := 2.0
	
	for i := 1; i <= 50; i++ {
		m := (l + r) / 2.0
		if m * m < 2 {
			l = m
		}else {
			r = m
		}
	
		fmt.Printf("Step %d %.15f -> %.15f\n", i, sm, m)
		sm = m
	}
}

func main() {
	fmt.Println("Newton")
	newton()

	fmt.Println("binary search")
	binSearch()
}

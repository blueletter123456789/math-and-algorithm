package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

// type Point struct {
// 	X, Y int
// }

// func main() {
// 	n := NumStdin()
// 	points := make([]Point, n)
// 	for i := 0; i < n; i++ {
// 		x := NumStdin()
// 		y := NumStdin()
// 		points[i] = Point{x, y}
// 	}

// 	sum_x, sum_y := 0, 0
// 	sort.SliceStable(points, func(i, j int) bool { 
// 		return points[i].X < points[j].X 
// 	})
	
// 	for i := 0; i < n; i++ {
// 		sum_x += points[i].X * (2*(i+1) - n - 1)
// 	}

// 	sort.SliceStable(points, func(i, j int) bool {
// 		return points[i].Y < points[j].Y
// 	})

// 	for i := 0; i < n; i++ {
// 		sum_y += points[i].Y * (2*(i+1) - n - 1)
// 	}

// 	fmt.Println(sum_x + sum_y)

// }
func main() {
	n := NumStdin()
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = NumStdin()
		y[i] = NumStdin()
	}

	sort.Ints(x)
	sort.Ints(y)

	sum_x, sum_y := 0, 0
	for i := 0; i < n; i++ {
		sum_x += x[i] * (2*(i+1) - n - 1)
		sum_y += y[i] * (2*(i+1) - n - 1)
	}

	fmt.Println(sum_x + sum_y)

}


func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

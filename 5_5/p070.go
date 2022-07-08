package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

type Point struct {
	X, Y int
}

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	k := NumStdin()
	area := make([]Point, n)
	for i := 0; i < n; i++ {
		area[i].X = NumStdin()
		area[i].Y = NumStdin()
	}

	ans := math.MaxInt64

	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					xl := area[a].X
					xr := area[b].X
					yb := area[c].Y
					yu := area[d].Y
					if is_over(xl, xr, yb, yu, k, area) {
						ans = min(ans, abs(yu-yb)*abs(xr-xl))
					}
				}
			}
		}
	}

	fmt.Println(ans)

}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func is_over(xl, xr, yb, yu, k int, lst[]Point) bool {
	cnt := 0
	for _, p := range lst {
		if xl <= p.X && p.X <= xr && yb <= p.Y && p.Y <= yu {
			cnt++
		}
	}

	return cnt >= k
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

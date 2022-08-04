package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	n := NumStdin()
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = NumStdin(), NumStdin()
	}
	
	a, b := NumStdin(), NumStdin()

	// 探索
	cnt := 0
	for i := 0; i < n; i++ {
		xa, ya := x[i]-a, y[i]-b
		xb, yb := x[(i+1)%n]-a, y[(i+1)%n]-b
		if ya > yb {
			xa, xb = xb, xa
			ya, yb = yb, ya
		}
		if ya <= 0 && 0 < yb && (xa*yb - xb*ya) < 0 {
			cnt++
		}
	}

	if cnt % 2 == 1 {
		fmt.Println("INSIDE")
	} else {
		fmt.Println("OUTSIDE")
	}
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

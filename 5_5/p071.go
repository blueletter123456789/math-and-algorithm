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

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	lines := make([][]float64, n)
	for i := 0; i < n; i++ {
		a := float64(NumStdin())
		b := float64(NumStdin())
		c := float64(NumStdin())
		lines[i] = []float64{a, b, c}
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if lines[i][0]*lines[j][1] == lines[j][0]*lines[i][1] {
				continue
			}
			x, y := point(lines[i][0],lines[i][1],lines[i][2],
				lines[j][0],lines[j][1],lines[j][2])
			if is_under(x-(1e-8), y-(1e-8), lines) {
				ans = math.Max(ans, x+y)
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

func point(a1, b1, c1, a2, b2, c2 float64) (x, y float64) {
	x = (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
	y = (c1*a2 - c2*a1) / (b1*a2 - b2*a1)
	return
}

func is_under(x, y float64, lines [][]float64) bool {
	for _, p := range lines {
		if (p[0]*x + p[1]*y) > p[2] {
			return false
		}
	}
	return true
}

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
	maxBufSize = 10000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

// 全探索
func main() {
	scan.Split(bufio.ScanWords)

	n := NumStdin()
	xLst := make([]int, n)
	yLst := make([]int, n)
	for i := 0; i < n; i++ {
		xLst[i] = NumStdin()
		yLst[i] = NumStdin()
	}

	ans := math.Pow10(6)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			tgt := math.Sqrt(
				math.Pow(float64(xLst[i]-xLst[j]), 2)+
				math.Pow(float64(yLst[i]-yLst[j]), 2))
			if ans > tgt {
				ans = tgt
			}
		}
	}
	fmt.Println(ans)
}

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
	M = 998244353
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	a := NumStdin()
	b := NumStdin()
	c := NumStdin()

	sum_a := ((a * (a+1)) / 2) % M
	sum_b := ((b * (b+1)) / 2) % M
	sum_c := ((c * (c+1)) / 2) % M

	// a, b, cを辺とする立方体について考える
	fmt.Println((((sum_a*sum_b)%M)*sum_c)%M)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

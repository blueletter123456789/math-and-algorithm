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

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	n := NumStdin()
	A := make([]int, n)
	for i := 1; i < n; i++ {
		A[i] = A[i-1] + NumStdin()
	}

	dist := 0
	m := NumStdin()
	prev := NumStdin()
	for i := 0; i < m-1; i++ {
		current := NumStdin()
		dist += abs(A[current-1] - A[prev-1])
		prev = current
	}

	fmt.Println(dist)
}

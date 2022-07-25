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
	maxBufSize = 1000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = NumStdin()
	}
	for i := 0; i < n; i++ {
		B[i] = NumStdin()
	}

	sort.Slice(A, func(i, j int) bool {return A[i] < A[j]})
	sort.Slice(B, func(i, j int) bool {return B[i] < B[j]})

	ans := 0
	for i := 0; i < n; i++ {
		ans += abs(A[i] - B[i])
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

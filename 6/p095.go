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
	A := make([]int, n+1)
	B := make([]int, n+1)

	for i := 1; i <= n; i++ {
		c := NumStdin()
		scoreA := 0
		scoreB := 0

		if c == 1 {
			scoreA = NumStdin()
		} else {
			scoreB = NumStdin()
		}
		A[i] = A[i-1] + scoreA
		B[i] = B[i-1] + scoreB
	}

	q := NumStdin()
	for i := 0; i < q; i++ {
		l, r := NumStdin(), NumStdin()
		fmt.Println(A[r]-A[l-1], B[r]-B[l-1])
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

func NumsStdin(n int) []int {
	scan.Scan()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

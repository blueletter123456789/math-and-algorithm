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
	maxBufSize = 100000000000
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

func NumsStdin(n int) []int {
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = NumStdin()
	}
	return nums
}

func main() {
	n := NumStdin()
	q := NumStdin()
	lst := NumsStdin(n)

	cum := make([]int, n+1)
	for i := 0; i < n; i++ {
		cum[i+1] = cum[i] + lst[i]
	}

	for i := 0; i < q; i++ {
		l := NumStdin()
		r := NumStdin()
		fmt.Println(cum[r]-cum[l-1])
	}
}

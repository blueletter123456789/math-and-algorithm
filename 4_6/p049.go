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
	M = 1000000007
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	fmt.Println(fibo(n))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func fibo(n int) int {
	lst := make([]int, n)
	lst[0], lst[1] = 1, 1
	for i := 2; i < n; i++ {
		lst[i] = (lst[i-2] + lst[i-1]) % M
	}
	return lst[n-1]
}

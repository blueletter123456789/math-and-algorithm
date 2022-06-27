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

	ret := pow(n-1)
	ans := ret[2][0]*2 + ret[2][1] + ret[2][2]
	fmt.Println(ans % M)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func pow(n int) [3][3]int {
	A := [3][3]int{{1, 1, 1}, {1, 0, 0}, {0, 1, 0}}
	ret := [3][3]int{}

	flg := true
	for i := 0; i< 60; i++ {
		if (n & (1 << i)) != 0 {
			if flg {
				ret = A
				flg = false
			} else {
				ret = multi(ret, A)
			}
		}
		A = multi(A, A)
	}

	return ret
}

func multi(a, b [3][3]int) [3][3]int {
	ret := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				ret[i][j] += a[i][k] * b[k][j]
				ret[i][j] %= M
			}
		}
	}

	return ret
}

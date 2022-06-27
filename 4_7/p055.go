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

	ans := pow(n-1)
	fmt.Println((ans[1][0] + ans[1][1]) % M)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func pow(n int) [2][2]int {
	A := [2][2]int{{2, 1}, {1, 0}}
	ret := [2][2]int{}
	flg := true
	for i := 0; i < 60; i++ {
		if (n & (1 << i)) != 0 {
			if flg {
				ret = A
				flg = false
			}else{
				ret = multi(ret, A)
			}
		}
		A = multi(A, A)
	}
	return ret
}

func multi(a, b [2][2]int) [2][2]int {
	ret := [2][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				ret[i][j] += a[i][k] * b[k][j]
				ret[i][j] %= M
			}
		}
	}
	return ret
}

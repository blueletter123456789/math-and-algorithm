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
	// ans := (((n * (n+1)) % M) * modpow(2, M-2)) % M
	ans := ((n * (n+1)) / 2) % M

	fmt.Println((ans * ans) % M)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

// func modpow(a, r int) int {
// 	p := a
// 	ans := 1
// 	for i := 0; i < 30; i++ {
// 		if (r & (1 << i)) != 0 {
// 			ans = (ans * p) % M
// 		}
// 		p = (p * p) % M
// 	}
// 	return ans
// }

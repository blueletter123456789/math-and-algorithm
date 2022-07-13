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

	factor := make([]int, n)
	factor[0] = 1
	for i := 1; i < n; i++ {
		factor[i] = (factor[i-1]*i) % M
	}

	ans := 0
	for i := 0; i < n; i++ {
		num := NumStdin()
		x := ncr(n-1, i, factor)
		ans += num*x
		ans %= M
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

func ncr(a, b int, factor []int) int {
	return (factor[a] * power((factor[a-b]*factor[b]) % M, M-2)) % M
}

func power(a, r int) int {
	p := a
	res := 1
	for i := 0; i < 60; i++ {
		if r & (1 << i) != 0 {
			res *= p
			res %= M
		}
		p *= p
		p %= M
	}
	return res
}

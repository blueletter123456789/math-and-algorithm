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
	x := NumStdin()
	y := NumStdin()

	fmt.Println(ncr(x+y, y))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func ncr(x, y int) int {
	lim := max(x, y)
	fact := make([]int, lim+1)
	fact[0] = 1
	for i := 1; i <= lim; i++ {
		fact[i] = (fact[i-1]*i)%M
	}
	return (fact[x] * pow((fact[x-y]*fact[y])%M, M-2)) % M
}

func max(a, b int) int {
	if a < b{
		return b
	}
	return a
}

func pow(a, b int) int {
	p := a
	ans := 1

	for i := 0; i <= 30; i++ {
		if b & (1 << i) != 0 {
			ans *= p
			ans %= M
		}
		p *= p
		p %= M
	}

	return ans
}

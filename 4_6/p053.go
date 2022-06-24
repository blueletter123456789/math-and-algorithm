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

	sum := pow(4, n+1)-1
	fact := pow(3, M-2)
	fmt.Println((sum * fact)%M)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func pow(x, r int) int {
	p := x
	res := 1
	for i := 0; i <= 65; i++ {
		if (r & (1 << i)) != 0 {
			res *= p
			res %= M
		}
		p *= p
		p %= M
	}

	return res
}

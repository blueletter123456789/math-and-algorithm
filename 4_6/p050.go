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
	a := NumStdin()
	b := NumStdin()

	fmt.Println(pow(a, b))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
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

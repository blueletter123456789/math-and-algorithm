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
	maxBufSize = 1000000
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

func erathosthenes(n int) []bool {
	prime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if prime[i]{
			for j := i*2; j <= n; j+=i {
				prime[j] = false
			}
		}
	}
	return prime
}

func main() {
	n := NumStdin()
	primes := erathosthenes(n)

	for i := 0; i <= n; i++ {
		if primes[i] {
			fmt.Println(i)
		}
	}
}

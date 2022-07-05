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
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	k := NumStdin()

	ans := 0
	for a := 1; a <= n; a++ {
		for b := max(1, a-k+1); b <= min(n, k+a-1); b++ {
			for c := max(1, a-k+1); c <= min(n, k+a-1); c++ {
				if abs(b-c) < k {
					ans++
				}
			}
		}
	}

	total := n * n * n

	fmt.Println(total - ans)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

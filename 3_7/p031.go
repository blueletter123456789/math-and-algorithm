package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 1000000000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func NumsStdin(n int) []int {
	scan.Scan()
	nums := make([]int, n)

	for i, v := range strings.Split(scan.Text(), " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	n := NumStdin()
	lst := NumsStdin(n)

	dp := make([]int, n+1)
	dp[1] = lst[0]
	for i := 1; i < n; i++ {
		dp[i+1] = max(dp[i], dp[i-1]+lst[i])
	}
	
	fmt.Println(dp[n])
}

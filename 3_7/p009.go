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

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	A := NumsStdin(n)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= A[i] && dp[i][j-A[i]]{
				dp[i+1][j] = true
			}
		}
	}
	if dp[n][s] {
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}

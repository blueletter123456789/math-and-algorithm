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

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(x, y int) int {
	if x <= y{
		return x
	}
	return y
}

func main() {
	n := NumStdin()
	h := NumsStdin(n)
	dp := make([]int, n)
	
	dp[1] = absInt(h[0]-h[1])
	for i := 2; i < n; i++ {
		dp[i] = minInt(dp[i-2]+absInt(h[i]-h[i-2]), 
				dp[i-1]+absInt(h[i]-h[i-1]))
	}
	
	fmt.Println(dp[n-1])
}

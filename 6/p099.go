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

var n int
var visited []bool
var dp []int
var G [][]int

func main() {
	n = NumStdin()
	G = make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := NumStdin(), NumStdin()
		G[a-1] = append(G[a-1], b-1)
		G[b-1] = append(G[b-1], a-1)
	}
	visited = make([]bool, n)
	dp = make([]int, n)

	dfs(0)

	ans := 0
	for i := 0; i < n; i++ {
		ans += dp[i] * (n - dp[i])
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

func dfs(u int) {
	visited[u] = true
	dp[u] = 1
	for _, v := range G[u] {
		if visited[v] {
			continue
		}
		dfs(v)
		dp[u] += dp[v]
	}
}

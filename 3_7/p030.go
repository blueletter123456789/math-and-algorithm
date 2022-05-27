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
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func chmax(a *int, b int){
	if *a < b{
		*a = b
	}
}

func main() {
	scan.Split(bufio.ScanWords)
	n := NumStdin()
	w := NumStdin()

	// Structでデータ構造を作成したほうがよかった
	wLst := make([]int, n)
	vLst := make([]int, n)

	// 二次元配列作成は実装での手段が正しいかも
	dp := make([][]int, n+1)
	dp[0] = make([]int, w+1)
	for i := 0; i < n; i++ {
		wLst[i] = NumStdin()
		vLst[i] = NumStdin()
		dp[i+1] = make([]int, w+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			// 逆ループだと一次元配列で済んだ
			if j < wLst[i] {
				dp[i+1][j] = dp[i][j]
				continue
			}
			chmax(&dp[i+1][j], dp[i][j])
			chmax(&dp[i+1][j], dp[i][j-wLst[i]]+vLst[i])
		}
	}

	fmt.Println(dp[n][w])
}

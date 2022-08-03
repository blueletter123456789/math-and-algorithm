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
	T := make([]int, n)
	sumT := 0
	for i := 0; i < n; i++ {
		T[i] = NumStdin()
		sumT += T[i]
	}

	dp := make([][]bool, n+1)
	dp[0] = make([]bool, sumT/2+1)
	dp[0][0] = true

	for i := 0; i < n; i++ {
		dp[i+1] = make([]bool, sumT/2+1)
		for j := 0; j <= sumT/2; j++ {
			dp[i+1][j] = dp[i][j]
			if T[i] > j {
				continue
			}
			dp[i+1][j] = dp[i+1][j] || dp[i][j-T[i]]
		}
	}
	
	ans := 0
	for i, flg := range dp[n] {
		if flg {
			ans = i
		}
	}
	fmt.Println(sumT - ans)

	// 確認用コード
	// minAns := 1 << 60
	// for i := 0; i < (1 << n); i++ {
	// 	ans1, ans2 := 0, 0
	// 	for j := 0; j < n; j++ {
	// 		if i & (1 << j) != 0 {
	// 			ans1 += T[j]
	// 		} else {
	// 			ans2 += T[j]
	// 		}
	// 	}
	// 	minAns = min(minAns, max(ans1, ans2))
	// }
	// fmt.Println(minAns)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

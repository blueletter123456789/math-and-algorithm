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

	facts := make([]int, n+1)
	facts[0] = 1
	for i := 1; i <= n; i++ {
		facts[i] = (facts[i-1]*i) % M
	}

	// 長さiでループ
	for i := 1; i <= n; i++ {
		ans := 0
		// 個数jでループ
		// 長さiで割った時に繰り上げた個数分
		for j := 1; j <= (n+i-1)/i; j++ {
			// (i-1)×(j-1)分は同一の囲い
			ans += ncr(n - (i-1) * (j-1), j, facts)
			ans %= M
		}

		fmt.Println(ans)
	}
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func ncr(n, r int, facts []int) int {
	return division(facts[n], (facts[r]*facts[n-r])%M)
}

func division(a, b int) int {
	return (a * modpow(b, M-2)) % M
}

func modpow(a, r int) int {
	p := a
	ret := 1
	for i := 0; i < 30; i++ {
		if (r & (1 << i)) != 0 {
			ret *= p
			ret %= M
		}
		p *= p
		p %= M
	}
	return ret
}

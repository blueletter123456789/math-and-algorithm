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
	l, r := NumStdin(), NumStdin()

	/*
	n < 5*10^5までの素数
	l <= i <= rにてi/各素数で更新
	*/
	primes := make([]bool, r-l+1)
	if l == 1 {
		primes[0] = true
	}

	for i := 2; i*i <= r; i++ {
		// 区間内のiの最小の倍数
		// 最小の倍数の場合は次の倍数から更新
		bottom := ((l+i-1)/i)*i
		if i == bottom {
			bottom *= 2
		}

		for j := bottom; j <= r; j+=i {
			primes[j-l] = true
		}
	}

	cnt := 0
	for _, val := range primes {
		if !val {
			cnt++
		}
	}
	fmt.Println(cnt)

	// 全探索
	// ans := 0
	// for i := l; i <= r; i++ {
	// 	if isPrime(i) {
	// 		ans++
	// 	}
	// }
	// fmt.Println(ans)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

// func isPrime(n int) bool {
// 	if n == 1 {
// 		return false
// 	}
// 	ret := true
// 	for i := 2; i*i <= n; i++ {
// 		if n % i == 0 {
// 			ret = false
// 		}
// 	}
// 	return ret
// }

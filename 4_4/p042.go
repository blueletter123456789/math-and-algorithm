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

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func factor_cnt(n int) []int {
	nums := make([]int, n+1)

	nums[1] = 1
	for i := 2; i <= n; i++ {
		// 1と自身のカウントするため
		nums[i] = 2
	}

	// numsの初期化をするループを実行せず、i:=1で開始しても可
	for i := 2; i <= n; i++ {
		for j := 2*i; j <= n; j += i {
			nums[j] += 1
		}
	}
	return nums
}

func main() {
	n := NumStdin()
	factors := factor_cnt(n)
	ans := 0
	for i, v := range factors {
		ans += i*v
	}
	fmt.Println(ans)
}


/* Using　Summation
func main(){
	n := NumStdin()
	ans := 0
	var cnt, sum int
	for i := 1; i <= n; i++ {
		cnt = n / i
		sum = (cnt*(cnt+1))/2
		ans += i*sum
	}
	fmt.Println(ans)
}
*/


/* 全探索
func main(){
	fmt.Println(factors)

	primes := prime(n)

	ans := 0
	for i := 1; i <= n; i++ {
		if !primes[i] {
			ans += i*divisor(i)
		}else{
			ans += 2*i
		}
	}

	fmt.Println(ans)
}

func prime(n int) []bool {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i]{
			for j := i*2; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	return primes
}

func divisor(n int) int {
	res := 0
	for i := 1; i*i <= n; i++ {
		if n % i == 0 {
			res += 1
			if i != n/i {
				res += 1
			}
		}
	}
	return res
}
*/

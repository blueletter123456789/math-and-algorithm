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
	maxBufSize = 10000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

var fm_cand map[int]struct{} = make(map[int]struct{}, 400000)

func main() {
	// 11桁までの短調増加の数の積を計算
	fm_cand = funcM(0, 0)

	n := NumStdin()
	b := NumStdin()

	ans := 0
	for fm := range fm_cand {
		// m = f(m) + Bとなるmを計算
		m := fm + b
		// 本来のf(m)を計算
		prod_m := product(m)

		// 問題文中の条件と一致するか
		// prod_m == fmと同義
		if m - prod_m == b && m <= n {
			ans++
		}
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

func funcM(digit, m int) map[int]struct{} {

	if digit == 11 {
		return map[int]struct{}{product(m): {}}
	}

	min_val := m % 10
	ret := map[int]struct{}{}
	
	for i := min_val; i <= 9; i++ {
	
		r := funcM(digit+1, m * 10 + i)
	
		for key := range r {
			ret[key] = struct{}{}
		}
	}
	
	return ret
}

func product(m int) int {
	if m == 0 {
		return 0
	} else {
		ans := 1

		for m >= 1 {
			ans *= m % 10
			m /= 10
		}
		return ans
	}
}

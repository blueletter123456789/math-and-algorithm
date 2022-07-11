package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	nums := make([]int, n)
	pows := make([]int, n+1)
	pows[0] = 1
	for i := 0; i < n; i++ {
		nums[i] = NumStdin()
		pows[i+1] = (pows[i]*2) % M
	}

	sort.Ints(nums)
	ans := 0
	for i := 0; i < n; i++ {
		ans += (nums[i] * pows[i]) % M
		ans %= M
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

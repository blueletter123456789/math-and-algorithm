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
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = NumStdin()
	}
	sort.Ints(nums)

	ans := 0
	for i := 0; i < n; i++ {
		ans += nums[i] * ((i+1)*2 - n - 1)
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

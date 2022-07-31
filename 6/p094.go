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
	
	prev := NumStdin()
	ans := prev
	for i := 1; i < n-1; i++ {
		cur := NumStdin()
		ans += min(prev, cur)
		prev = cur
	}
	fmt.Println(ans + prev)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
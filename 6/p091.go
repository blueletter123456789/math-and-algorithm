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
	x := NumStdin()

	cnt := 0
	for a := 1; a <= n; a++ {
		for b := a+1; b <= n; b++ {
			c := x - a - b
			if c <= a || c <= b || c > n {
				continue
			}
			cnt++
		}
	}

	fmt.Println(cnt)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func IntStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func FloatsStdin(n int) []float64 {
	scan.Scan()
	nums := make([]float64, n)
	for i, str_num := range strings.Split(scan.Text(), " ") {
		num, err := strconv.ParseFloat(str_num, 64)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func main() {
	n := IntStdin()
	bLst := FloatsStdin(n)
	rLst := FloatsStdin(n)
	var b, r float64
	for i := 0; i < n; i++ {
		b += bLst[i] / float64(n)
		r += rLst[i] / float64(n)
	}
	fmt.Println(b + r)
}

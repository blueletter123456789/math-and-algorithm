package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	n := IntStdin()
	total := 0.0
	for i := 1; i <= n; i++ {
		total += float64(n) / float64(i)
	}
	fmt.Println(total)
}

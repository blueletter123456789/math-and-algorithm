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
	a := NumStdin()
	b := NumStdin()
	c := NumStdin()
	d := NumStdin()

	ans1 := a*c
	ans2 := a*d
	ans3 := b*c
	ans4 := b*d

	fmt.Println(max(max(ans1, ans2), max(ans3, ans4)))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
